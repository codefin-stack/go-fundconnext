package fundconnext

import (
	"encoding/json"
	"fmt"
	"io"
)

// BankAccount Verified V4
type BankAccount struct {
	BankCode         string  `json:"bankCode"`
	BankBranchCode   *string `json:"bankBranchCode"`
	BankAccountNo    string  `json:"bankAccountNo"`
	Default          bool    `json:"default"`
	FinnetCustomerNo *string `json:"finnetCustomerNo"`
	Currency         string  `json:"currency,omitempty"`
}

type IndividualAccountDocument struct {
	IdentificationCardType   string        `json:"identificationCardType"`
	PassportCountry          *string       `json:"passportCountry"`
	CardNumber               string        `json:"cardNumber"`
	AccountId                string        `json:"accountId"`
	IcLicense                string        `json:"icLicense"`
	AccountOpenDate          string        `json:"accountOpenDate"`
	MailingAddressSameAsFlag *string       `json:"mailingAddressSameAsFlag"`
	Mailing                  *Address      `json:"mailing"`
	MailingMethod            string        `json:"mailingMethod"`
	InvestmentObjective      string        `json:"investmentObjective"`
	InvestmentObjectiveOther *string       `json:"investmentObjectiveOther"`
	RedemptionBankAccounts   []BankAccount `json:"redemptionBankAccounts"`
	SubscriptionBankAccounts []BankAccount `json:"subscriptionBankAccounts"`
	Approved                 bool          `json:"approved"`
	OpenOmnibusFormFlag      *bool         `json:"openOmnibusFormFlag"`
}

type IndividualAccountFile struct {
	IdentificationCardType string  `json:"identificationCardType"`
	PassportCountry        *string `json:"passportCountry"`
	CardNumber             string  `json:"cardNumber"`
	AccountId              string  `json:"accountId"`
	Approved               bool    `json:"approved"`
}

func (f *FundConnext) CreateIndividualAccount(individualAccDoc IndividualAccountDocument) error {
	url := "/api/customer/individual/account/v4"
	body, err := json.Marshal(individualAccDoc)
	if err != nil {
		f.cfg.Logger.Error("[Func CreateIndividualAccount] Error json.Marshal ::", err)
		return err
	}
	_, err = f.APICall("POST", url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func CreateIndividualAccount] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdateIndividualAccount(individualAccDoc IndividualAccountDocument) error {
	url := "/api/customer/individual/account/v4"
	body, err := json.Marshal(individualAccDoc)
	if err != nil {
		f.cfg.Logger.Error("[Func UpdateIndividualAccount] Error json.Marshal ::", err)
		return err
	}
	_, err = f.APICall("PUT",url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func UpdateIndividualAccount] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

// TODO:
// func (f *FundConnext) UploadIndividualAccountFile(fileType string, individualAccountFile IndividualAccountFile) error {
// 	cfg := MakeAPICallerConfig(f)
// 	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
// 	body, err := json.Marshal(individualAccountFile)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = CallFCAPI(f.token, "PUT", url, body, cfg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (f *FundConnext) UploadIndividualAccountFile(fileType string, body io.Reader) error {
	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
	_, err := f.APICallFormData("POST", url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func UploadIndividualAccountFile] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}
