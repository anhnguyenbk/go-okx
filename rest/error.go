package rest

import "fmt"

type OKXError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewOKXError(code, message, details string) OKXError {
	return OKXError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

var _ error = (*OKXError)(nil)

func (e OKXError) Error() string {
	return fmt.Sprintf("code: %s, message: %s, details: %s", e.Code, e.Message, e.Details)
}
