package fundconnext

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type APICallerConfig struct {
	Timeout     *time.Duration
	ContentType string
	Logger      *log.Logger
	Proxy       string
}

var (
	DEMO_URL       = "https://lh-mock-api.codefin.dev/lh"
	STAGING_URL    = "https://stage.fundconnext.com"
	PRODUCTION_URL = "https://www.fundconnext.com"
	MOCK_URL       = "http://mock-api:8085/lh"
)

type FCAuthentication struct {
	AccessToken string
	SACode      string
	SAId        string
	Username    string
	IssuedAt    int64
	ExpiresAt   int64
}

type FundConnext struct {
	In             IFundConnext
	cfg            *FCConfiguration
	token          string
	authentication *FCAuthentication
}

type IFundConnext interface {
	APICall(method, url string, req interface{}) ([]byte, error)
}

type FCConfiguration struct {
	Username string
	Password string
	Timeout  time.Duration
	Env      string
	Logger   *log.Logger
	Proxy    string
}

func ToOptStr(str string) *string {
	r := &str
	return r
}

func ToOptBool(flag bool) *bool {
	r := &flag
	return r
}

func ToOptFloat(num float32) *float32 {
	r := &num
	return r
}

func ToOptFloat64(num float64) *float64 {
	r := &num
	return r
}

func ToOptInt(num int) *int {
	r := &num
	return r
}

func (f *FundConnext) getUrl() string {
	switch f.cfg.Env {
	case "staging":
		return STAGING_URL
	case "production":
		return PRODUCTION_URL
	case "mock-api":
		return MOCK_URL
	default:
		return DEMO_URL
	}
}

func MakeAPICallerConfig(f *FundConnext) *APICallerConfig {
	return &APICallerConfig{
		Timeout: &f.cfg.Timeout,
		Logger:  f.cfg.Logger,
		Proxy:   f.cfg.Proxy,
	}
}

func (f *FundConnext) reTokenize() error {
	url := f.getUrl()
	token, err := Login(url, f.cfg.Username, f.cfg.Password, f.cfg.Proxy)
	if err != nil {
		return err
	}

	f.authentication = &FCAuthentication{
		AccessToken: token.AccessToken,
		SACode:      token.SACode,
		SAId:        token.Claims.SellingAgentId,
		Username:    token.Claims.Username,
		ExpiresAt:   token.Claims.ExpiresAt,
		IssuedAt:    token.Claims.IssuedAt,
	}
	return nil
}

func (f *FundConnext) APICall(method, url string, req interface{}) ([]byte, error) {
	if f.authentication == nil || !time.Now().Before(time.Unix(f.authentication.ExpiresAt, 0).Add(time.Minute*-15)) {
		f.cfg.Logger.Infoln("[Func APICall] Info Re-tokenized")
		if err := f.reTokenize(); err != nil {
			return nil, err
		}
	}

	env := f.getUrl()
	cfg := MakeAPICallerConfig(f)
	cfg.ContentType = "application/json"
	resp, err := CallFCAPI(env, f.authentication.AccessToken, method, url, req, cfg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FundConnext) APICallFormData(method, url string, req interface{}) ([]byte, error) {
	if f.authentication == nil || !time.Now().Before(time.Unix(f.authentication.ExpiresAt, 0).Add(time.Minute*-15)) {
		f.cfg.Logger.Infoln("[Func APICall] Info Re-tokenized")
		if err := f.reTokenize(); err != nil {
			return nil, err
		}
	}

	env := f.getUrl()
	cfg := MakeAPICallerConfig(f)
	cfg.ContentType = "multipart/form-data"
	resp, err := CallFCAPI(env, f.authentication.AccessToken, method, url, req, cfg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FundConnext) Configure(cfg *FCConfiguration) {
	if cfg.Logger == nil {
		cfg.Logger = log.New()
	}
	f.cfg = cfg
}

func New(cfg *FCConfiguration) *FundConnext {
	fc := new(FundConnext)
	fc.Configure(cfg)
	fc.In = fc
	return fc
}
