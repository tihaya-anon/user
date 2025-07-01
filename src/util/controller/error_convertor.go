package controller_util

import (
	"MVC_DI/global/model"
	"MVC_DI/vo/resp"
	"slices"
)

func ExposeError(response *resp.TResponse, err error, status ...any) *resp.TResponse {
	appErr, ok := err.(*model.AppError)
	if !ok {
		return response.SystemError()
	}
	if slices.Contains(status, appErr.StatusKey) {
		return response.CustomerError().WithData(*appErr)
	}
	return response.SystemError()
}
