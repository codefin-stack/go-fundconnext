package fundconnext

var (
	UnAuthorizedError     = MakeFCError("E000", "Unauthorized access")
	UnProcessableEntry    = MakeFCError("EXXX", "Unprocessable entry")
	LtfBalanceNotFound    = MakeFCError("E339", "LTF Balance Not Found")
	UnSupportAdvanceOrder = MakeFCError("E274", "This fund does not support advance order. Please change effective date.")
)

type ErrMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FCError struct {
	ErrMsg ErrMsg `json:"errMsg"`
}

func (e *FCError) Error() string {
	return e.ErrMsg.Message
}

func (e *FCError) Code() string {
	return e.ErrMsg.Code
}

func MakeInternalError(message string) error {
	return &FCError{
		ErrMsg: ErrMsg{
			Code:    "I000",
			Message: message,
		},
	}
}
func MakeFCError(code, message string) error {
	return &FCError{
		ErrMsg: ErrMsg{
			Code:    code,
			Message: message,
		},
	}
}

func IsFCError(err error) *FCError {
	fcErr, ok := err.(*FCError)
	if ok {
		return fcErr
	}
	return nil
}

func Is(err error, errorType interface{}) *FCError {
	typeError, ok := errorType.(error)
	if ok {
		fcErrType := IsFCError(typeError)
		if fcErrType != nil {
			fcErr := IsFCError(err)
			if fcErr != nil {
				if fcErrType.Code() == "EXXX" || fcErr.Code() == fcErrType.Code() {
					return fcErr
				}
			}
		}
	}
	return nil
}
