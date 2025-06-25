package global_model

import "fmt"

type AppError struct {
	Code    string
	Message *string
	Detail  any
}

func (e *AppError) Error() string {
	if e.Message == nil {
		return fmt.Sprintf("[%s]", e.Code)
	}
	return fmt.Sprintf("[%s] %s", e.Code, *e.Message)
}

func NewAppError() *AppError {
	return &AppError{}
}

func (e *AppError) WithCode(code string) *AppError {
	e.Code = code
	return e
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = &message
	return e
}

func (e *AppError) WithDetail(detail any) *AppError {
	e.Detail = detail
	return e
}
