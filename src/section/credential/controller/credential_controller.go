package credential_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	credential_service "MVC_DI/section/credential/service"
	"MVC_DI/vo/resp"
)

type CredentialController struct {
	CredentialService credential_service.CredentialService
	Logger *logrus.Logger
}

func (ctrl *CredentialController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `credential`")
}