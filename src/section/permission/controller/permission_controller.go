package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"MVC_DI/section/permission/service"
	"MVC_DI/vo/resp"
)

type PermissionController struct {
	PermissionService service.PermissionService
	Logger *logrus.Logger
}

func (ctrl *PermissionController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `permission`")
}