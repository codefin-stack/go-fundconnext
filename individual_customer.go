package fundconnext

import (
	"encoding/json"
	"fmt"
	"io"
)

// SpouseDocument Verified V4
type SpouseDocument struct {
	THFirstName *string `json:"thFirstName"`
	THLastName  *string `json:"thLastName"`
	ENFirstName *string `json:"enFirstName"`
	ENLastName  *string `json:"enLastName"`
}

// SuitabilityForm Verified V4
type SuitabilityForm struct {
	SuitNo1  *int   `json:"suitNo1"`
	SuitNo2  *int   `json:"suitNo2"`
	SuitNo3  *int   `json:"suitNo3"`
	SuitNo4  *[]int `json:"suitNo4"`
	SuitNo5  *int   `json:"suitNo5"`
	SuitNo6  *int   `json:"suitNo6"`
	SuitNo7  *int   `json:"suitNo7"`
	SuitNo8  *int   `json:"suitNo8"`
	SuitNo9  *int   `json:"suitNo9"`
	SuitNo10 *int   `json:"suitNo10"`
	SuitNo11 *int   `json:"suitNo11"`
	SuitNo12 *int   `json:"suitNo12"`
}

// Address Verified V4
type Address struct {
	No          string  `json:"no"`
	Floor       *string `json:"floor"`
	Building    *string `json:"building"`
	RoomNo      *string `json:"roomNo"`
	Soi         *string `json:"soi"`
	Road        *string `json:"road"`
	Moo         *string `json:"moo"`
	SubDistrict string  `json:"subDistrict"`
	District    string  `json:"district"`
	Province    string  `json:"province"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
}

type PartialIndividualCustomerDocument struct {
	IdentificationCardType    string  `json:"identificationCardType"`
	PassportCountry           *string `json:"passportCountry"`
	CardNumber                string  `json:"cardNumber"`
	SuitabilityRiskLevel      *string `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate *string `json:"suitabilityEvaluationDate"`
	Fatca                     *bool   `json:"fatca"`
	FatcaDeclarationDate      *string `json:"fatcaDeclarationDate"`
	CddScore                  *string `json:"cddScore"`
	CddDate                   *string `json:"cddDate"`
	ReferralPerson            *string `json:"referralPerson"`
	Approved                  bool    `json:"approved"`
}

type IndividualCustomerDocument struct {
	IdentificationCardType         string           `json:"identificationCardType"`
	PassportCountry                *string          `json:"passportCountry"`
	CardNumber                     string           `json:"cardNumber"`
	CardExpiryDate                 string           `json:"cardExpiryDate"`
	AccompanyingDocument           *string          `json:"accompanyingDocument"`
	Title                          string           `json:"title"`
	TitleOther                     *string          `json:"titleOther"`
	ENFirstName                    string           `json:"enFirstName"`
	ENLastName                     string           `json:"enLastName"`
	THFirstName                    string           `json:"thFirstName"`
	THLastName                     string           `json:"thLastName"`
	BirthDate                      string           `json:"birthDate"`
	Nationality                    string           `json:"nationality"`
	MobileNumber                   string           `json:"mobileNumber"`
	Email                          *string          `json:"email"`
	Phone                          *string          `json:"phone"`
	Fax                            *string          `json:"fax"`
	MaritalStatus                  string           `json:"maritalStatus"`
	Spouse                         *SpouseDocument  `json:"spouse"`
	OccupationId                   string           `json:"occupationId"`
	OccupationOther                *string          `json:"occupationOther"`
	BusinessTypeId                 *string          `json:"businessTypeId"`
	BusinessTypeOther              *string          `json:"businessTypeOther"`
	MonthlyIncomeLevel             string           `json:"monthlyIncomeLevel"`
	AssetValue                     *string          `json:"assetValue"`
	IncomeSource                   string           `json:"incomeSource"`
	IncomeSourceOther              *string          `json:"incomeSourceOther"`
	IdentificationDocument         Address          `json:"identificationDocument"`
	CurrentAddressSameAsFlag       *string          `json:"currentAddressSameAsFlag"`
	Current                        *Address         `json:"current"`
	CompanyName                    *string          `json:"companyName"`
	Work                           *Address         `json:"work"`
	WorkPosition                   *string          `json:"workPosition"`
	RelatedPoliticalPerson         *bool            `json:"relatedPoliticalPerson"`
	PoliticalRelatedPersonPosition *string          `json:"politicalRelatedPersonPosition"`
	CanAcceptFxRisk                bool             `json:"canAcceptFxRisk"`
	CanAcceptDerivativeInvestment  bool             `json:"canAcceptDerivativeInvestment"`
	SuitabilityRiskLevel           string           `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate      string           `json:"suitabilityEvaluationDate"`
	Fatca                          bool             `json:"fatca"`
	FatcaDeclarationDate           string           `json:"fatcaDeclarationDate"`
	CDDScore                       *string          `json:"cddScore"`
	CDDDate                        *string          `json:"cddDate"`
	ReferralPerson                 *string          `json:"referralPerson"`
	ApplicationDate                string           `json:"applicationDate"`
	IncomeSourceCountry            string           `json:"incomeSourceCountry"`
	AcceptBy                       *string          `json:"acceptBy"`
	OpenFundConnextFormFlag        string           `json:"openFundConnextFormFlag"`
	Approved                       bool             `json:"approved"`
	VulnerableFlag                 *bool            `json:"vulnerableFlag"`
	VulnerableDetail               *string          `json:"vulnerableDetail"`
	NDIDFlag                       *bool            `json:"ndidFlag"`
	NDIDRequestId                  *string          `json:"ndidRequestId"`
	OpenChannel                    *string          `json:"openChannel"`
	SuitabilityForm                *SuitabilityForm `json:"suitabilityForm"`
	InvestorClass                  *string          `json:"investorClass"`
}

type IndividualCustomerFile struct {
	IdentificationCardType string  `json:"identificationCardType"`
	PassportCountry        *string `json:"passportCountry"`
	CardNumber             string  `json:"cardNumber"`
	Approved               bool    `json:"approved"`
}

func (f *FundConnext) CreateIndividualCustomer(identificationDoc IndividualCustomerDocument) error {
	url := "/api/customer/individual/v4"
	body, err := json.Marshal(identificationDoc)

	if err != nil {
		f.cfg.Logger.Error("[Func CreateIndividualCustomer] Error json.Marshal ::", err)
		return err
	}

	_, err = f.APICall("POST", url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func CreateIndividualCustomer] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdateIndividualCustomer(identificationDoc IndividualCustomerDocument) error {
	url := "/api/customer/individual/v4"
	body, err := json.Marshal(identificationDoc)

	if err != nil {
		f.cfg.Logger.Error("[Func UpdateIndividualCustomer] Error json.Marshal ::", err)
		return err
	}

	_, err = f.APICall("PUT", url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func UpdateIndividualCustomer] Error CallToFundConnext ::", err)
		return err
	}

	return nil
}

func (f *FundConnext) UpdatePartialIndividualCustomer(partialIndividualCustomerDocument PartialIndividualCustomerDocument) error {
	url := "/api/customer/individual"
	body, err := json.Marshal(partialIndividualCustomerDocument)
	if err != nil {
		f.cfg.Logger.Error("[Func UpdatePartialIndividualCustomer] Error json.Marshal ::", err)
		return err
	}
	_, err = f.APICall("PATCH", url, body)
	if err != nil {
		f.cfg.Logger.Error("[Func UpdatePartialIndividualCustomer] Error CallToFundConnext ::", err)
		return err
	}
	return nil
}

// TODO:
// func (f *FundConnext) UploadIndividualCustomerFile(fileType string, individualCustomerFile IndividualCustomerFile) error {
// 	cfg := MakeAPICallerConfig(f)
// 	url := fmt.Sprintf("/api/customer/individual/account/%s/upload", fileType)
// 	body, err := json.Marshal(individualCustomerFile)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = CallFCAPI(f.token, "PUT", url, body, cfg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (f *FundConnext) UploadIndividualCustomerFile(fileType string, body io.Reader) error {
	url := fmt.Sprintf("/api/customer/individual/%s/upload", fileType)
	_, err := f.APICallFormData("POST", url, body)
	if err != nil {
		return err
	}
	return nil
}
