package fundconnext

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/sanchaimac/go-amnet/data"

	"github.com/mitchellh/mapstructure"
)

type Download struct {
	DataType    data.FundConnextFileType
	Header      DownloadHeader
	Body        []interface{}
	HeaderBytes []byte
	BodyBytes   []byte
}

// DownloadHeader ...
type DownloadHeader struct {
	Date    string
	SA      string
	Records int
	Version string
}

// New ...
func (Download) New(fileName string) *Download {
	return &Download{
		DataType: data.FundConnextFileTypeMapping[fileName],
	}
}

// Scan input will be the struct structure that you want output to be look like
func (d *Download) Scan(requestedStruct interface{}) error {
	return mapstructure.Decode(d.Body, &requestedStruct)
}

// Download FundConnext file (can save multiple paths)
func (f *FundConnext) Download(asOfDate string, fileType data.FundConnextFileType, optionalSavePath ...string) (result Download, err error) {
	url := fmt.Sprintf("/api/files/%s/%s.zip", asOfDate, fileType.String())
	// out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg) out, err := f.APICall("GET", url, make([]byte, 0))
	out, err := f.APICall("GET", url, make([]byte, 0))
	if err != nil {
		return result, err
	}

	saveErr := saveFile(out, optionalSavePath)
	if saveErr != nil {
		return Download{}, saveErr
	}

	zipReader, err := zip.NewReader(bytes.NewReader(out), int64(len(out)))
	if err != nil {
		return Download{}, err
	}

	if zipReader == nil || zipReader.File[0] == nil {
		return result, errors.New("empty zip file")
	}

	zipFile := zipReader.File[0]
	return parseDownloadFile(zipFile, fileType)
}

// Private Function only accessible for download func
func saveFile(data []byte, fullPath []string) error {

	if len(data) <= 0 || len(fullPath) <= 0 {
		return nil
	}

	for _, path := range fullPath {
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		out, err := os.Create(path)
		if err != nil {
			return err
		}

		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, bytes.NewBuffer(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func parseDownloadFile(zipFile *zip.File, fileType data.FundConnextFileType) (Download, error) {
	f, err := zipFile.Open()
	if err != nil {
		return Download{}, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	// Get File Header
	scanner.Scan()
	headerString := scanner.Text()
	Header := strings.Split(headerString, "|")
	headerBytes := []byte(headerString)
	var DataStruct []interface{}
	var bodyBytes []byte

	for scanner.Scan() {
		lineString := scanner.Text()
		bodyBytes = append(bodyBytes, []byte(lineString)...)
		lineData := strings.Split(lineString, "|")
		downloadType := reflect.TypeOf(fileType.ModelType())
		reflectVal := reflect.New(downloadType)

		for key, value := range lineData {
			if key >= reflectVal.Elem().NumField() {
				continue
			}
			if reflectVal.Elem().Field(key).CanSet() {
				field := reflectVal.Elem().Field(key)
				tag := downloadType.Field(key).Tag
				var nullable bool

				if tagValue, ok := tag.Lookup("fundconnext"); ok {
					if tagValue == "nullable" {
						nullable = true
					}
				}
				t, err := toReflectValue(value, field.Interface(), nullable)
				if err != nil {
					return Download{}, err
				}

				reflectVal.Elem().Field(key).Set(t)
			}
		}

		DataStruct = append(DataStruct, reflectVal.Elem().Interface())
	}
	var asOfDate string
	p0 := fileType.Header().AsOfDate
	if p0 >= 0 {
		asOfDate = Header[p0]
	}
	var saCode string
	p1 := fileType.Header().SACode
	if p1 >= 0 {
		saCode = Header[p1]
	}
	var records int
	p2 := fileType.Header().TotalRecord
	if p2 >= 0 {
		records = StringToInt(Header[p2])
	}

	var version string
	if len(Header) >= 4 {
		version = Header[3]
	}
	return Download{
		DataType: fileType,
		Header: DownloadHeader{
			Date:    asOfDate,
			SA:      saCode,
			Records: records,
			Version: version,
		},
		Body:        DataStruct,
		HeaderBytes: headerBytes,
		BodyBytes:   bodyBytes,
	}, nil
}

func toReflectValue(text string, value interface{}, nullable bool) (reflect.Value, error) {
	var err error
	switch value.(type) {
	case *float32:
		var r *float32
		if text != "" {
			s, err := strconv.ParseFloat(text, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]float32{float32(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]float32{0}[0]
		}
		return reflect.ValueOf(r), nil
	case float32:
		var r float32
		if text != "" {
			s, err := strconv.ParseFloat(text, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = float32(s)
		}
		return reflect.ValueOf(r), nil
	case *float64:
		var r *float64
		if text != "" {
			s, err := strconv.ParseFloat(text, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]float64{s}[0]
		}
		if !nullable && r == nil {
			r = &[]float64{0}[0]
		}
		return reflect.ValueOf(r), nil
	case float64:
		var s float64
		if text != "" {
			s, err = strconv.ParseFloat(text, 64)

			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(s), nil
	case *int8:
		var r *int8
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]int8{int8(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]int8{0}[0]
		}
		return reflect.ValueOf(r), nil
	case int8:
		var rd int8
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int8(s)
		}
		return reflect.ValueOf(rd), nil
	case *int16:
		var r *int16
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]int16{int16(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]int16{0}[0]
		}
		return reflect.ValueOf(r), nil
	case int16:
		var rd int16
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int16(s)
		}
		return reflect.ValueOf(rd), nil
	case *int32:
		var r *int32
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]int32{int32(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]int32{0}[0]
		}
		return reflect.ValueOf(r), nil
	case int32:
		var rd int32
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int32(s)
		}
		return reflect.ValueOf(rd), nil
	case *int64:
		var r *int64
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]int64{s}[0]
		}
		if !nullable && r == nil {
			r = &[]int64{0}[0]
		}
		return reflect.ValueOf(r), nil
	case int64:
		var rd int64
		if text != "" {
			rd, err = strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(rd), nil
	case *int:
		var r *int
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]int{int(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]int{0}[0]
		}
		return reflect.ValueOf(r), nil
	case int:
		var rd int
		if text != "" {
			s, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = int(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint8:
		var r *uint8
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]uint8{uint8(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]uint8{0}[0]
		}
		return reflect.ValueOf(r), nil
	case uint8:
		var rd uint8
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 8)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint8(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint16:
		var r *uint16
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]uint16{uint16(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]uint16{0}[0]
		}
		return reflect.ValueOf(r), nil
	case uint16:
		var rd uint16
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 16)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint16(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint32:
		var r *uint32
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]uint32{uint32(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]uint32{0}[0]
		}
		return reflect.ValueOf(r), nil
	case uint32:
		var rd uint32
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 32)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint32(s)
		}
		return reflect.ValueOf(rd), nil
	case *uint64:
		var r *uint64
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]uint64{s}[0]
		}
		if !nullable && r == nil {
			r = &[]uint64{0}[0]
		}
		return reflect.ValueOf(r), nil
	case uint64:
		var rd uint64
		if text != "" {
			rd, err = strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
		}
		return reflect.ValueOf(rd), nil
	case *uint:
		var r *uint
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			r = &[]uint{uint(s)}[0]
		}
		if !nullable && r == nil {
			r = &[]uint{0}[0]
		}
		return reflect.ValueOf(r), nil
	case uint:
		var rd uint
		if text != "" {
			s, err := strconv.ParseUint(text, 10, 64)
			if err != nil {
				return reflect.ValueOf(nil), errors.New("type conversion failed")
			}
			rd = uint(s)
		}
		return reflect.ValueOf(rd), nil
	case bool:
		var r bool
		if text != "" {
			switch text {
			case "Y", "y", "1", "T", "t":
				r = true
			case "N", "n", "0", "F", "f":
				r = false
			default:
				r = true
			}
		}
		return reflect.ValueOf(r), nil
	case *bool:
		var r bool
		if text != "" {
			switch text {
			case "Y", "y", "1", "T", "t":
				r = true
			case "N", "n", "0", "F", "f":
				r = false
			default:
				r = true
			}
			return reflect.ValueOf(&r), nil
		} else {
			return reflect.ValueOf(nil), nil
		}
	case *string:
		if nullable && text == "" {
			return reflect.ValueOf(nil), nil
		}
		return reflect.ValueOf(&text), nil
	case string:
		return reflect.ValueOf(text), nil
	default:
		return reflect.ValueOf(nil), errors.New("no type conversion support")
	}
}
