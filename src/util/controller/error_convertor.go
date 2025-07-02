package controller_util

import (
	"MVC_DI/global/model"
	"MVC_DI/vo/resp"
	"slices"

	"github.com/sirupsen/logrus"
)

func ExposeError(response *resp.TResponse, logger *logrus.Logger, err error, status ...any) *resp.TResponse {
	appErr, ok := err.(*model.AppError)
	if !ok {
		logger.Errorf("failed to convert to app error: %v", err)
		return response.SystemError()
	}
	if slices.Contains(status, appErr.StatusKey) {
		logger.Warnf("exposed error: %v", appErr)
		return response.CustomerError().WithData(*appErr)
	}
	logger.Errorf("hidden error: %v", appErr)
	return response.SystemError()
}
