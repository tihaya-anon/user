package resp

import (
	"MVC_DI/global/enum"
	"MVC_DI/vo/resp/common"
)

type TResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func NewResponse() *TResponse {
	return &TResponse{}
}
func (response *TResponse) Success() *TResponse {
	response.Code = enum.CODE.SUCCESS
	response.Msg = enum.MSG.SUCCESS
	return response
}

func (response *TResponse) SuccessWithData(data any) *TResponse {
	*response = *response.Success()
	response.Data = data
	return response
}

func (response *TResponse) Error(code string, msg string) *TResponse {
	response.Code = code
	response.Msg = msg
	return response
}

func (response *TResponse) SystemError() *TResponse {
	return response.Error(enum.CODE.SYSTEM_ERROR, enum.MSG.SYSTEM_ERROR)
}

func (response *TResponse) CustomerError(msg string) *TResponse {
	return response.Error(enum.CODE.CUSTOMER_ERROR, msg)
}

func (response *TResponse) ThirdPartyError(msg string) *TResponse {
	return response.Error(enum.CODE.THIRD_PARTY_ERROR, msg)
}

func (response *TResponse) ValidationError(error *common.ValidationError) *TResponse {
	response.Code = enum.CODE.VALIDATION_ERROR
	response.Msg = enum.MSG.VALIDATION_ERROR
	response.Data = &error
	return response
}

func (response *TResponse) AllArgsConstructor(code string, msg string, data any) *TResponse {
	response.Code = code
	response.Msg = msg
	response.Data = data
	return response
}
