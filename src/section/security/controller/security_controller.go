package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"MVC_DI/section/security/service"
	"MVC_DI/vo/resp"
)

type SecurityController struct {
	SecurityService service.SecurityService
	Logger *logrus.Logger
}

func (ctrl *SecurityController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `security`")
}