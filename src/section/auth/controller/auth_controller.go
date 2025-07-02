package controller

import (
	"MVC_DI/gen/api"
	global_enum "MVC_DI/global/enum"
	"MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	"MVC_DI/section/auth/service"
	"MVC_DI/security"
	controller_uitl "MVC_DI/util/controller"
	"MVC_DI/vo/resp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	AuthService service.AuthService
	Logger      *logrus.Logger
}

func (ctrl *AuthController) LoginUser(ctx *gin.Context) *resp.TResponse {
	response := resp.NewResponse()

	userLoginRequest, validationError := controller_uitl.BindValidation[api.UserLoginRequest](ctx)
	if validationError != nil {
		return nil
	}

	userLoginDto := dto.UserLoginDto{
		Secret:     userLoginRequest.Secret,
		Identifier: userLoginRequest.Identifier,
		Type:       string(userLoginRequest.Type),
	}
	userLoginRespDto, err := ctrl.AuthService.LoginUser(ctx, userLoginDto)
	if err != nil {
		return controller_uitl.ExposeError(response, ctrl.Logger, err,
			auth_enum.UNKNOWN_CREDENTIAL{},
			auth_enum.PASSWORD_WRONG{},
			auth_enum.EMAIL_CODE_WRONG{},
			auth_enum.GOOGLE_2FA_WRONG{},
			auth_enum.OAUTH_WRONG{},
		)
	}
	// TODO expire config
	security.SetSessionId(ctx, userLoginRespDto.SessionId, 3600, "/", "", true, true)
	authSessionResponse := api.AuthSessionResponse{Token: &userLoginRespDto.Token}
	return response.Success().WithData(authSessionResponse)
}

func (ctrl *AuthController) LogoutUser(ctx *gin.Context) *resp.TResponse {
	response := resp.NewResponse()

	sessionId := security.GetSessionId(ctx)
	if sessionId == nil {
		return response.AllArgsConstructor(global_enum.CODE_MISSING_TOKEN, global_enum.MSG_MISSING_TOKEN, nil)
	}

	err := ctrl.AuthService.LogoutUser(ctx, *sessionId)
	if err != nil {
		return controller_uitl.ExposeError(response, ctrl.Logger, err,
			auth_enum.UNKNOWN_SESSION{},
		)
	}
	return response.Success()
}
