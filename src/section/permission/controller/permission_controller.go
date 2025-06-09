package permission_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	permission_service "MVC_DI/section/permission/service"
	"MVC_DI/vo/resp"
)

type PermissionController struct {
	PermissionService permission_service.PermissionService
	Logger *logrus.Logger
}

func (ctrl *PermissionController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `permission`")
}