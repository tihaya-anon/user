package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"MVC_DI/section/credential/service"
	"MVC_DI/vo/resp"
)

type CredentialController struct {
	CredentialService service.CredentialService
	Logger *logrus.Logger
}

func (ctrl *CredentialController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `credential`")
}