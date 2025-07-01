package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"MVC_DI/section/user/service"
	"MVC_DI/vo/resp"
)

type UserController struct {
	UserService service.UserService
	Logger *logrus.Logger
}

func (ctrl *UserController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `user`")
}