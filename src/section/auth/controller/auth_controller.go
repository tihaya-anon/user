package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/vo/resp"
)

type AuthController struct {
	AuthService auth_service.AuthService
	Logger *logrus.Logger
}

func (ctrl *AuthController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `auth`")
}