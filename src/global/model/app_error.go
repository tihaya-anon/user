package global_model

import (
	"MVC_DI/global/enum"
	"fmt"
)

type AppError struct {
	StatusKey any
	status    *enum.Status
	Detail    *any
}

func (e *AppError) Error() string {
	if e.Detail != nil {
		return fmt.Sprintf("[%s] %s: %v", e.status.Code(), e.status.Msg(), *e.Detail)
	}
	return fmt.Sprintf("[%s] %s", e.status.Code(), e.status.Msg())
}

func NewAppError() *AppError {
	return &AppError{}
}

// WithStatusKey
//
// must use struct enum from package global.enum
func (e *AppError) WithStatusKey(statusKey any) *AppError {
	status := enum.STATUS_MAP[statusKey]
	e.status = status
	return e
}

// WithStatusKey
//
// must use struct enum from package global.enum
func (e *AppError) WithStatusKeyOptionalMap(statusKey any, optionalMap *enum.StatusMap) *AppError {
	status := (*optionalMap)[statusKey]
	e.status = status
	return e
}
func (e *AppError) WithDetail(detail any) *AppError {
	e.Detail = &detail
	return e
}
