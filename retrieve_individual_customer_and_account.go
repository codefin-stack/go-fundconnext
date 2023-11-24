package fundconnext

import (
	"encoding/json"
	"fmt"
)

type CrsDetail struct {
	CountryOfTaxResidence string  `json:"countryOfTaxResidence"`
	Tin                   string  `json:"tin"`
	Reason                *string `json:"reason"`
	ReasonDesc            *string `json:"reasonDesc"`
}
type RetrievalIndividualCustomerProfileAndAccount struct {
	IdentificationCardType         string              `json:"identificationCardType"`
	PassportCountry                *string             `json:"passportCountry"`
	CardNumber                     string              `json:"cardNumber"`
	CardExpiryDate                 string              `json:"cardExpiryDate"`
	AccompanyingDocument           *string             `json:"accompanyingDocument"`
	Title                          string              `json:"title"`
	TitleOther                     *string             `json:"titleOther"`
	EnFirstName                    string              `json:"enFirstName"`
	EnLastName                     string              `json:"enLastName"`
	ThFirstName                    string              `json:"thFirstName"`
	ThLastName                     string              `json:"thLastName"`
	BirthDate                      string              `json:"birthDate"`
	Nationality                    string              `json:"nationality"`
	MobileNumber                   string              `json:"mobileNumber"`
	Email                          *string             `json:"email"`
	Phone                          *string             `json:"phone"`
	Fax                            *string             `json:"fax"`
	MaritalStatus                  string              `json:"maritalStatus"`
	Spouse                         *SpouseDocument     `json:"spouse"`
	OccupationId                   int                 `json:"occupationId"`
	OccupationOther                *string             `json:"occupationOther"`
	BusinessTypeId                 *int                `json:"businessTypeId"`
	BusinessTypeOther              *string             `json:"businessTypeOther"`
	MonthlyIncomeLevel             string              `json:"monthlyIncomeLevel"`
	AssetValue                     *float64            `json:"assetValue"`
	IncomeSource                   string              `json:"incomeSource"`
	IncomeSourceOther              *string             `json:"incomeSourceOther"`
	IdentificationDocument         Address             `json:"identificationDocument"`
	CurrentAddressSameAsFlag       *string             `json:"currentAddressSameAsFlag"`
	Current                        *Address            `json:"current"`
	CompanyName                    *string             `json:"companyName"`
	Work                           *Address            `json:"work"`
	WorkPosition                   *string             `json:"workPosition"`
	RelatedPoliticalPerson         *bool               `json:"relatedPoliticalPerson"`
	PoliticalRelatedPersonPosition *string             `json:"politicalRelatedPersonPosition"`
	CanAcceptFxRisk                bool                `json:"canAcceptFxRisk"`
	CanAcceptDerivativeInvestment  bool                `json:"canAcceptDerivativeInvestment"`
	SuitabilityRiskLevel           int                 `json:"suitabilityRiskLevel"`
	SuitabilityEvaluationDate      string              `json:"suitabilityEvaluationDate"`
	Fatca                          bool                `json:"fatca"`
	FatcaDeclarationDate           string              `json:"fatcaDeclarationDate"`
	CddScore                       *int                `json:"cddScore"`
	CddDate                        *string             `json:"cddDate"`
	ReferralPerson                 *string             `json:"referralPerson"`
	ApplicationDate                string              `json:"applicationDate"`
	IncomeSourceCountry            string              `json:"incomeSourceCountry"`
	AcceptBy                       *string             `json:"acceptedBy"`
	OpenFundConnextFormFlag        string              `json:"openFundConnextFormFlag"`
	Approved                       bool                `json:"approved"`
	VulnerableFlag                 *bool               `json:"vulnerableFlag"`
	VulnerableDetail               *string             `json:"vulnerableDetail"`
	NdidFlag                       *bool               `json:"ndidFlag"`
	NdidRequestId                  *string             `json:"ndidRequestId"`
	OpenChannel                    *string             `json:"openChannel"`
	SuitabilityForm                *SuitabilityForm    `json:"suitabilityForm"`
	InvestorClass                  *string             `json:"investorClass"`
	Accounts                       *[]RetrievalAccount `json:"accounts"`
	ApprovedDate                   *string             `json:"approvedDate"`

	// add from version 5
	CrsTaxResidenceInCountriesOtherThanTheUs bool        `json:"crsTaxResidenceInCountriesOtherThanTheUS"`
	CrsPlaceOfBirthCountry                   string      `json:"crsPlaceOfBirthCountry"`
	CrsPlaceOfBirthCity                      string      `json:"crsPlaceOfBirthCity"`
	CrsDeclarationDate                       string      `json:"crsDeclarationDate"`
	CrsDetails                               []CrsDetail `json:"crsDetails"`
}

type RetrievalAccount struct {
	AccountId                string         `json:"accountId"`
	IcLicense                string         `json:"icLicense"`
	AccountOpenDate          string         `json:"accountOpenDate"`
	MailingAddressSameAsFlag *string        `json:"mailingAddressSameAsFlag"`
	Mailing                  *Address       `json:"mailing"`
	MailingMethod            string         `json:"mailingMethod"`
	InvestmentObjective      string         `json:"investmentObjective"`
	InvestmentObjectiveOther *string        `json:"investmentObjectiveOther"`
	RedemptionBankAccounts   []BankAccount  `json:"redemptionBankAccounts"`
	SubscriptionBankAccounts *[]BankAccount `json:"subscriptionBankAccounts"`
	ApprovedDate             *string        `json:"approvedDate"`
	OpenOmnibusFormFlag      *bool          `json:"openOmnibusFormFlag"`
}

func (f *FundConnext) RetrieveIndividualCustomerProfileAndAccount(cardNumber string) (*RetrievalIndividualCustomerProfileAndAccount, error) {
	url := fmt.Sprintf("/api/customer/individual/investor/profile/v5?cardNumber=%s", cardNumber)

	out, err := f.APICall("GET", url, make([]byte, 0))
	if err != nil {
		f.cfg.Logger.Error("[Func RetrieveIndividualCustomerProfileAndAccount] Error CallToFundConnext ::", err)
		return nil, err
	}

	var results *RetrievalIndividualCustomerProfileAndAccount
	if err := json.Unmarshal(out, &results); err != nil {
		f.cfg.Logger.Error("[Func RetrieveIndividualCustomerProfileAndAccount] Error json.Marshal ::", err)
		return nil, err
	}
	return results, nil
}
