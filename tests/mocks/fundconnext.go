package mock

import (
	"log"

	"github.com/sanchaimac/go-amnet"
	"github.com/sirupsen/logrus"
)

type FundConnext struct {
	*fundconnext.FundConnext
	Mockf MockFundConnext
}

type MockFundConnext struct {
	APICall *ExpectedAPICall
}

type ExpectedAPICall struct {
	Return []byte
	Error  error
}

func NewFundConnext(m MockFundConnext) *FundConnext {
	mock := &FundConnext{
		Mockf: m,
		FundConnext: fundconnext.New(&fundconnext.FCConfiguration{
			Logger: logrus.New(),
		}),
	}

	mock.In = mock
	return mock
}

func (f *FundConnext) APICall(method, url string, req interface{}) ([]byte, error) {
	if f.Mockf.APICall != nil {
		log.Println("Called Mock APICall")
		return f.Mockf.APICall.Return, f.Mockf.APICall.Error
	}
	return f.FundConnext.APICall(method, url, req)
}
