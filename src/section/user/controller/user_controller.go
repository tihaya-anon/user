package user_controller

import (
	user_service "MVC_DI/section/user/service"
	"MVC_DI/vo/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	UserService user_service.UserService
	Logger      *logrus.Logger
}

func (ctrl *UserController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().Success().WithData("hello `user`")
}
