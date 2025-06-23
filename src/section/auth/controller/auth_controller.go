package auth_controller

import (
	"MVC_DI/gen/api"
	auth_dto "MVC_DI/section/auth/dto"
	auth_service "MVC_DI/section/auth/service"
	controller_uitl "MVC_DI/util/controller"
	"MVC_DI/vo/resp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	AuthService auth_service.AuthService
	Logger      *logrus.Logger
}

func (ctrl *AuthController) Hello(ctx *gin.Context) *resp.TResponse {
	return resp.NewResponse().SuccessWithData("hello `auth`")
}

func (ctrl *AuthController) LoginUser(ctx *gin.Context) *resp.TResponse {
	userLoginRequest, validationError := controller_uitl.BindValidation[api.UserLoginRequest](ctx)
	response := resp.NewResponse()
	if validationError != nil {
		return response.ValidationError(validationError)
	}
	userLoginDto := auth_dto.UserLoginDto{
		Secret:     userLoginRequest.Secret,
		Identifier: userLoginRequest.Identifier,
		Type:       string(userLoginRequest.Type),
	}
	userLoginRespDto, err := ctrl.AuthService.LoginUser(ctx, userLoginDto)
	if err != nil {
		return response.SystemError(err)
	}
	authSessionResponse := api.AuthSessionResponse{SessionId: &userLoginRespDto.SessionId, Token: &userLoginRespDto.Token}
	return response.SuccessWithData(authSessionResponse)
}
