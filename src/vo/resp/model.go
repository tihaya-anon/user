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

func (response *TResponse) Error(code string, error error) *TResponse {
	var msg string
	if error != nil {
		msg = error.Error()
	} else {
		msg = enum.MSG.SYSTEM_ERROR
	}
	response.Code = code
	response.Msg = msg
	return response
}

func (response *TResponse) SystemError(error error) *TResponse {
	return response.Error(enum.CODE.SYSTEM_ERROR, error)
}

func (response *TResponse) CustomerError(error error) *TResponse {
	return response.Error(enum.CODE.CUSTOMER_ERROR, error)
}

func (response *TResponse) ThirdPartyError(error error) *TResponse {
	return response.Error(enum.CODE.THIRD_PARTY_ERROR, error)
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
