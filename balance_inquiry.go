package fundconnext

import (
	"context"
	"encoding/json"

	"github.com/machinebox/graphql"
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

func (f *FundConnext) BalanceInquiry(accountId string, req *graphql.Request) (*BalanceInquiryResults, error) {
	f.cfg.Logger.Infoln("[Funconnext:BalanceInquiry] AccountId: ", accountId)

	f.cfg.Logger.Warning("[AmnetAPI:BalanceInquiry] Call check balance inquiry to Amet!")
	url := "api/account/balances"
	// out, err := f.APICall("GET", url, make([]byte, 0))
	// if err != nil {
	// 	return nil, err
	// }
	out, err := f.APICallAmnet(context.Background(), url, req)
	if err != nil {
		return nil, err
	}

	// Convert the interface to a map
	data := out.(map[string]interface{})

	// Convert the map to JSON bytes again
	jsonBytes, _ := json.Marshal(data)
	var results *BalanceInquiryResults
	json.Unmarshal([]byte(jsonBytes), &results)
	return results, nil
}
