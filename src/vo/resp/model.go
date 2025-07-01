package resp

import (
	"MVC_DI/global/enum"
	"MVC_DI/vo/resp/common"
)

type TResponse struct {
	Code enum.Code `json:"code"`
	Msg  enum.Msg  `json:"msg,omitempty"`
	Data any       `json:"data,omitempty"`
}

func NewResponse() *TResponse {
	return &TResponse{}
}
func (response *TResponse) Success() *TResponse {
	return response.Status(enum.SUCCESS{})
}

func (response *TResponse) WithData(data any) *TResponse {
	response.Data = data
	return response
}
// Status
//
// must use struct enum from package global.enum
func (response *TResponse) Status(statusKey any) *TResponse {
	status := enum.STATUS_MAP[statusKey]
	response.Code = status.Code()
	response.Msg = status.Msg()
	return response
}
func (response *TResponse) Error(code enum.Code, msg enum.Msg) *TResponse {
	response.Code = code
	response.Msg = msg
	return response
}

func (response *TResponse) SystemError() *TResponse {
	return response.Status(enum.SYSTEM_ERROR{})
}

func (response *TResponse) CustomerError() *TResponse {
	return response.Status(enum.CUSTOMER_ERROR{})
}

func (response *TResponse) ThirdPartyError() *TResponse {
	return response.Status(enum.THIRD_PARTY_ERROR{})
}

func (response *TResponse) ValidationError(error *common.ValidationError) *TResponse {
	response.Code = enum.CODE_VALIDATION_ERROR
	response.Msg = enum.MSG_VALIDATION_ERROR
	response.Data = &error
	return response
}

func (response *TResponse) AllArgsConstructor(code enum.Code, msg enum.Msg, data any) *TResponse {
	response.Code = code
	response.Msg = msg
	response.Data = data
	return response
}
