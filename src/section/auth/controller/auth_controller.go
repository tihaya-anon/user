package auth_controller

import (
	"MVC_DI/gen/api"
	"MVC_DI/global/enum"
	auth_dto "MVC_DI/section/auth/dto"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/security"
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
	response := resp.NewResponse()

	userLoginRequest, validationError := controller_uitl.BindValidation[api.UserLoginRequest](ctx)
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

	security.SetSessionId(ctx, userLoginRespDto.SessionId, 3600, "/", "", true, true)
	authSessionResponse := api.AuthSessionResponse{Token: &userLoginRespDto.Token}
	return response.SuccessWithData(authSessionResponse)
}

func (ctrl *AuthController) LogoutUser(ctx *gin.Context) *resp.TResponse {
	response := resp.NewResponse()

	sessionId := security.GetSessionId(ctx)
	if sessionId == nil {
		return response.AllArgsConstructor(enum.CODE.MISSING_TOKEN, enum.MSG.MISSING_TOKEN, nil)
	}

	err := ctrl.AuthService.LogoutUser(ctx, *sessionId)
	if err != nil {
		return response.SystemError(err)
	}

	return response.Success()
}
