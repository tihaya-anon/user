package credential_controller

import (
	credential_service "MVC_DI/section/credential/service"
	"MVC_DI/vo/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CredentialController struct {
	CredentialService credential_service.CredentialService
	Logger            *logrus.Logger
}

func (ctrl *CredentialController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `credential`")
}
