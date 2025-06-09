package test_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	test_service "MVC_DI/section/test/service"
	"MVC_DI/vo/resp"
)

type TestAController struct {
	TestAService test_service.TestAService
	Logger *logrus.Logger
}

func (ctrl *TestAController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `test-a`")
}