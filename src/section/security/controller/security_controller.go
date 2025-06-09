package security_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	security_service "MVC_DI/section/security/service"
	"MVC_DI/vo/resp"
)

type SecurityController struct {
	SecurityService security_service.SecurityService
	Logger *logrus.Logger
}

func (ctrl *SecurityController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `security`")
}