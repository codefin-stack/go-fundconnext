package tests

import (
	"archive/zip"
	"bytes"
	"strings"
	"testing"

	"github.com/codefin-stack/go-fundconnext/data"
	mock "github.com/codefin-stack/go-fundconnext/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMockZip(fileContent string) []byte {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	f, err := w.Create("test.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(fileContent))
	if err != nil {
		panic(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func TestDownloadBankAccountUnitholderData(t *testing.T) {
	// old format (7 columns)
	oldContent := `20201010|SA01|1|1.0
ACC01|AMC01|UH01|S|BBL|1234567890|Y`

	// new format (9 columns)
	newContent := `20201010|SA01|1|2.0
ACC01|AMC01|UH01|S|BBL|1234567890|Y|INDIVIDUAL|USD`

	tests := []struct {
		name     string
		content  string
		expected data.BankAccountUnitholderData
	}{
		{
			name:    "Older format",
			content: oldContent,
			expected: data.BankAccountUnitholderData{
				AccountID:              "ACC01",
				AMCCode:                "AMC01",
				UnitholderID:           "UH01",
				TransactionType:        "S",
				BankCode:               "BBL",
				BankAccountNumber:      "1234567890",
				DefaultFlagBankAccount: "Y",
				UnitholderType:         "",
				Currency:               "",
			},
		},
		{
			name:    "V2.0 format",
			content: newContent,
			expected: data.BankAccountUnitholderData{
				AccountID:              "ACC01",
				AMCCode:                "AMC01",
				UnitholderID:           "UH01",
				TransactionType:        "S",
				BankCode:               "BBL",
				BankAccountNumber:      "1234567890",
				DefaultFlagBankAccount: "Y",
				UnitholderType:         "INDIVIDUAL",
				Currency:               "USD",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zipBytes := createMockZip(tt.content)
			testTools := SetupTest(t, SetupOptions{
				Mock: &Mock{
					FundConnext: mock.MockFundConnext{
						APICall: &mock.ExpectedAPICall{
							Return: zipBytes,
							Error:  nil,
						},
					},
				},
			})
			defer testTools.Teardown(t)

			dl, err := testTools.FC.Download("20201010", data.BankAccountUnitholder)
			require.NoError(t, err)

			require.Len(t, dl.Body, 1)
			actual := dl.Body[0].(data.BankAccountUnitholderData)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDownloadUnitholderMappingData(t *testing.T) {
	// old format (4 columns)
	oldContent := `20201010|SA01|1|1.0
ACC01|AMC01|UH01|OMNIBUS`

	// new format (5 columns)
	newContent := `20201010|SA01|1|2.0
ACC01|AMC01|UH01|OMNIBUS|USD`

	tests := []struct {
		name     string
		content  string
		expected data.UnitholderMappingData
	}{
		{
			name:    "Older format",
			content: oldContent,
			expected: data.UnitholderMappingData{
				AccountID:    "ACC01",
				AMCCode:      "AMC01",
				UnitholderID: "UH01",
				AccountType:  "OMNIBUS",
				Currency:     "",
			},
		},
		{
			name:    "V2.0 format",
			content: newContent,
			expected: data.UnitholderMappingData{
				AccountID:    "ACC01",
				AMCCode:      "AMC01",
				UnitholderID: "UH01",
				AccountType:  "OMNIBUS",
				Currency:     "USD",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zipBytes := createMockZip(tt.content)
			testTools := SetupTest(t, SetupOptions{
				Mock: &Mock{
					FundConnext: mock.MockFundConnext{
						APICall: &mock.ExpectedAPICall{
							Return: zipBytes,
							Error:  nil,
						},
					},
				},
			})
			defer testTools.Teardown(t)

			dl, err := testTools.FC.Download("20201010", data.UnitholderMapping)
			require.NoError(t, err)

			require.Len(t, dl.Body, 1)
			actual := dl.Body[0].(data.UnitholderMappingData)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDownloadFundProfileCurrency(t *testing.T) {
	// real V4.0 row (100 columns) from the UAT FundProfile file of 20261022
	row := "LHDAILYUSD|LHFUND|กองทุนเปิด แอล เอช เดลี่ ยูเอสดี|LH DAILY USD FUND|F||N|N|20111201|3|N|Y|1400|100.00|0.00|1400|||||3||N|N|LHDAILYUSD|N|N|O|O|0|0|0||||P|Y|0|N|N||||1|R|||N|ธนาคารไทยพาณิชย์ จำกัด (มหาชน)|บริษัท หลักทรัพย์จัดการกองทุนแลนด์ แอนด์ เฮ้าส์ จำกัด|MF0641/2554||||||||USD|||||||||||||||||||||||||||||||||||||||||"
	cols := strings.Split(row, "|")
	require.Len(t, cols, 100)

	tests := []struct {
		name     string
		content  string
		currency string
	}{
		{name: "V4.0 format", content: "20261022|1|V4.0\n" + row, currency: "USD"},
		{name: "Older format (44 columns)", content: "20261022|1|V3.0\n" + strings.Join(cols[:44], "|"), currency: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zipBytes := createMockZip(tt.content)
			testTools := SetupTest(t, SetupOptions{
				Mock: &Mock{
					FundConnext: mock.MockFundConnext{
						APICall: &mock.ExpectedAPICall{
							Return: zipBytes,
							Error:  nil,
						},
					},
				},
			})
			defer testTools.Teardown(t)

			dl, err := testTools.FC.Download("20261022", data.FundProfile)
			require.NoError(t, err)

			require.Len(t, dl.Body, 1)
			actual := dl.Body[0].(data.FundProfileData)
			assert.Equal(t, "LHDAILYUSD", actual.FundCode)
			assert.Equal(t, tt.currency, actual.Currency)
		})
	}
}
