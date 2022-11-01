package fundconnext

import (
	"encoding/json"
	"fmt"
)

type LTFRedeemableUnitInquiryResponse struct {
	UnitholderId string  `json:"unitholderId"`
	FundCode     string  `json:"fundCode"`
	Unit         float64 `json:"unit"`
}

func (f *FundConnext) LTFRedeemableUnitInquiry(unitholderId string, fundCode string) (*LTFRedeemableUnitInquiryResponse, error) {
	url := fmt.Sprintf("/api/ltfBalances/redeemableUnit?unitholderId=%s&fundCode=%s", unitholderId, fundCode)
	f.cfg.Logger.Infoln("[Func LTFRedeemableUnitInquiry]: ", url)

	out, err := f.In.APICall("GET", url, make([]byte, 0))
	if err != nil {
		f.cfg.Logger.Infoln("[Func LTFRedeemableUnitInquiry] Call API Error:", err)
		return nil, err
	}
	f.cfg.Logger.Infoln("[Func LTFRedeemableUnitInquiry] Call API Result:", out)
	var results LTFRedeemableUnitInquiryResponse
	json.Unmarshal(out, &results)
	return &results, nil
}
