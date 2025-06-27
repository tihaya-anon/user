package controller_util

import (
	global_model "MVC_DI/global/model"
	"MVC_DI/vo/resp"
	"slices"
)

func ExposeError(response *resp.TResponse, err error, err_code ...string) *resp.TResponse {
	appErr, ok := err.(*global_model.AppError)
	if !ok {
		return response.SystemError()
	}
	if slices.Contains(err_code, appErr.Code) {
		return response.CustomerError(appErr.Error())
	}
	return response.SystemError()
}
