package fundconnext

import (
	"encoding/json"
	"fmt"
)

type RedemptionOrder struct {
	SaOrderReferenceNo    string   `json:"saOrderReferenceNo"`
	TransactionDateTime   string   `json:"transactionDateTime"`
	SACode                string   `json:"saCode"`
	AccountId             string   `json:"accountId"`
	UnitholderId          string   `json:"unitholderId"`
	FundCode              string   `json:"fundCode"`
	RedemptionType        string   `json:"redemptionType"`
	Amount                *float64 `json:"amount"`
	Unit                  *float64 `json:"unit"`
	EffectiveDate         string   `json:"effectiveDate"`
	PaymentType           string   `json:"paymentType"`
	BankCode              *string  `json:"bankCode"`
	BankAccount           *string  `json:"bankAccount"`
	Channel               string   `json:"channel"`
	ICLicense             string   `json:"icLicense"`
	BranchNo              *string  `json:"branchNo"`
	ForceEntry            string   `json:"forceEntry"`
	ReasonToSellLtfRmf    *string  `json:"reasonToSellLtfRmf"`
	RmfCapGainWhtChoice   *string  `json:"rmfCapGainWhtChoice"`
	RmfCapAmtRedeemChoice *string  `json:"RmfCapAmtRedeemChoice"`
	SellAllUnitFlag       *string  `json:"sellAllUnitFlag"`
	SettlementBankCode    *string  `json:"settlementBankCode"`
	SettlementBankAccount *string  `json:"settlementBankAccount"`
	Status                *string  `json:"status"`
	CollateralAccount     *string  `json:"collateralAccount"`
}

type RedemptionOrderResponse struct {
	TransactionId      string `json:"transactionId"`
	SaOrderReferenceNo string `json:"saOrderReferenceNo"`
	SettlementDate     string `json:"settlementDate"`
}

func (f *FundConnext) CreateRedemption(redemption RedemptionOrder) (*RedemptionOrderResponse, error) {
	url := "/api/redemptions/"
	body, err := json.Marshal(redemption)
	if err != nil {
		return nil, err
	}
	// resp, err := CallFCAPI(f.token, "POST", url, body, cfg)
	resp, err := f.In.APICall("POST", url, body)
	if err != nil {
		return nil, err
	}
	var results *RedemptionOrderResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) CancelRedemption(transactionId, force string) (*TransactionIDResponse, error) {
	url := fmt.Sprintf("/api/redemptions/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"force": force,
	})
	if err != nil {
		return nil, err
	}
	// resp, err := CallFCAPI(f.token, "DELETE", url, body, cfg)
	resp, err := f.APICall("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	var results *TransactionIDResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (f *FundConnext) ApproveRedemption(transactionId, status string) (*TransactionIDResponse, error) {
	url := fmt.Sprintf("/api/redemptions/%s", transactionId)
	body, err := json.Marshal(map[string]string{
		"status": status,
	})
	if err != nil {
		return nil, err
	}
	// resp, err := CallFCAPI(f.token, "PATCH", url, body, cfg)
	resp, err := f.APICall("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	var results *TransactionIDResponse
	err = json.Unmarshal(resp, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
