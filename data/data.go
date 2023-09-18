package data

// FundMappingData ...
type FundMappingData struct {
	AmcCode  string
	FundCode string
}

// FundProfileData ...
type FundProfileData struct {
	FundCode               string
	AMCCode                string
	FundNameTH             string
	FundNameEN             string
	FundPolicy             string
	TaxType                *string
	FIFFlag                string
	DividendFlag           string
	RegistrationDate       string
	FundRiskLevel          string
	FXRiskFlag             string
	FATCAAllowFlag         string
	BuyCutOffTime          string
	FstLowBuyVal           *float64
	NxtLowBuyVal           *float64
	SellCutOffTime         string
	LowSellVal             *float64 `fundconnext:"nullable"`
	LowSellUnit            *float64 `fundconnext:"nullable"`
	LowBalVal              *float64 `fundconnext:"nullable"`
	LowBalUnit             *float64 `fundconnext:"nullable"`
	SellSettlementDay      uint64
	SwitchingSettlementDay *uint64
	SwitchOutFlag          string
	SwitchInFlag           string
	FundClass              *string
	BuyPeriodFlag          string
	SellPeriodFlag         string
	SwitchInPeriodFlag     *string
	SwitchOutPeriodFlag    *string
	BuyPreOrderDay         float64
	SellPreOrderDay        float64
	SwitchPreOrderDay      float64
	AutoRedeemFund         *string
	BegIPODate             *string
	EndIPODate             *string
	PlainComplexFund       string
	DerivativesFlag        string
	LagAllocationDay       uint64
	SettlementHolidayFlag  string
	HealthInsurance        string
	PreviousFundCode       *string
	InvestorAlert          *string
	ISIN                   *string
	LowBalCondition        *string
}

// SwitchingMatrixData ...
type SwitchingMatrixData struct {
	FundCodeOut         string
	FundCodeIn          string
	SwitchSettlementDay int
	SwitchingType       *string
}

// FundHolidayData ...
type FundHolidayData struct {
	FundCode    string
	HolidayDate string
}

// TradeCalendarData ...
type TradeCalendarData struct {
	FundCode        string
	TransactionCode string
	TradeType       string
	TradeDate       string
}

// FeeData ...
type FeeData struct {
	EffectiveDate string
	FundCode      string
	FeeType       string
	FeeUnit       string
	MaximumFee    *float64
	ActualFee     *float64
	MinimumFee    *float64
	Remark        *string
	MaximumValue  *float64
	Filler1       *string
	Filler2       *string
	Filler3       *string
	Filler4       *string
	Filler5       *string
	Filler6       *string
	Filler7       *string
	Filler8       *string
	Filler9       *string
	Filler10      *string
	Filler11      *string
}

// FundPerformanceData ...
type FundPerformanceData struct {
	FundCode   string
	PYTDReturn *float64
	P3MReturn  *float64
	P6MReturn  *float64
	P1YReturn  *float64
	P3YReturn  *float64
	P5YReturn  *float64
	P10YReturn *float64
	PSIReturn  *float64
	P1YSD      *float64
	P3YSD      *float64
	P5YSD      *float64
	P10YSD     *float64
	NAVDate    string
}

// AccountProfileData ...
type AccountProfileData struct {
	SACode                              string
	AccountID                           string
	AccountOpenDate                     string
	ICLicense                           string
	BankCodeSubscription1               *string
	BankAccountSubscription1            *string
	DefaultFlagBankAccountSubscription1 *string
	BankCodeSubscription2               *string
	BankAccountSubscription2            *string
	DefaultFlagBankAccountSubscription2 *string
	BankCodeSubscription3               *string
	BankAccountSubscription3            *string
	DefaultFlagBankAccountSubscription3 *string
	BankCodeSubscription4               *string
	BankAccountSubscription4            *string
	DefaultFlagBankAccountSubscription4 *string
	BankCodeSubscription5               *string
	BankAccountSubscription5            *string
	DefaultFlagBankAccountSubscription5 *string
	BankCodeRedemption1                 *string
	BankAccountRedemption1              *string
	DefaultFlagBankAccountRedemption1   *string
	BankCodeRedemption2                 *string
	BankAccountRedemption2              *string
	DefaultFlagBankAccountRedemption2   *string
	BankCodeRedemption3                 *string
	BankAccountRedemption3              *string
	DefaultFlagBankAccountRedemption3   *string
	BankCodeRedemption4                 *string
	BankAccountRedemption4              *string
	DefaultFlagBankAccountRedemption4   *string
	BankCodeRedemption5                 *string
	BankAccountRedemption5              *string
	DefaultFlagBankAccountRedemption5   *string
	Gender                              *string
	Title                               *string
	FirstNameTH                         string
	LastNameTH                          *string
	FirstNameEN                         string
	LastNameEN                          *string
	SuitabilityTestDate                 string
	RiskProfileLevel                    string
	FXRiskFlag                          string
	FATCATestDate                       string
	FATCAFlag                           string
	VulnerableFlag                      *string
	OpenFundConnextFormFlag             *string
	OpenOmnibusFormFlag                 *string
	AccountStatus                       string
	ProcessStatus                       string
	LastUpdate                          string
	InvestorType                        string
}

// UnitholderMappingData ...
type UnitholderMappingData struct {
	AccountID    string
	AMCCode      string
	UnitholderID string
	AccountType  string
}

// BankAccountUnitholderData ...
type BankAccountUnitholderData struct {
	AccountID         string
	AMCCode           string
	UnitholderID      string
	TransactionType   string
	BankCode          string
	BankAccountNumber string
}
