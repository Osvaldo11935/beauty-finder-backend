package error_common

type CustomError struct {
	Code        string   `json:"code"`
	Message     string   `json:"message"`
	Description string   `json:"description"`
	Errs        []string `json:"errs"`
}

func NewCustomError(code string, message, description string, errs []string) *CustomError {
	return &CustomError{
		Code:        code,
		Errs:        errs,
		Message:     message,
		Description: description,
	}
}

func (e *CustomError) Error() string {
	return e.Description
}

func (e *CustomError) GetMessageError() string {
	return e.Message
}
func (e *CustomError) GetDescriptionErr() string {
	return e.Description
}
