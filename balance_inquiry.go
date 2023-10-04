package fundconnext

import (
	"encoding/json"
	"fmt"
)

type BalanceInquiry struct {
	UnitholderId  string  `json:"unitholderId"`
	FundCode      string  `json:"fundCode"`
	Unit          float64 `json:"unit"`
	Amount        float64 `json:"amount"`
	RemainUnit    float64 `json:"remainUnit"`
	RemainAmount  float64 `json:"remainAmount"`
	PendingAmount float64 `json:"pendingAmount"`
	PendingUnit   float64 `json:"pendingUnit"`
	AVGCost       float64 `json:"avgCost"`
	NAV           float64 `json:"nav"`
	NAVDate       string  `json:"navDate"`
}

type BalanceInquiryResults struct {
	Result []BalanceInquiry `json:"result"`
}

func (f *FundConnext) BalanceInquiry(accountNo string) (*BalanceInquiryResults, error) {
	f.cfg.Logger.Infoln("[Funconnext:BalanceInquiry] AccountNo: ", accountNo)
	url := fmt.Sprintf("/api/account/balances?accountNo=%s", accountNo)
	// out, err := CallFCAPI(f.token, "GET", url, make([]byte, 0), cfg)
	out, err := f.APICall("GET", url, make([]byte, 0))
	if err != nil {
		return nil, err
	}
	var results *BalanceInquiryResults
	json.Unmarshal(out, &results)
	return results, nil
}
