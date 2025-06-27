package auth_controller

import (
	"MVC_DI/gen/api"
	"MVC_DI/global/enum"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
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

func (ctrl *AuthController) LoginUser(ctx *gin.Context) *resp.TResponse {
	response := resp.NewResponse()

	userLoginRequest, validationError := controller_uitl.BindValidation[api.UserLoginRequest](ctx)
	if validationError != nil {
		return nil
	}

	userLoginDto := auth_dto.UserLoginDto{
		Secret:     userLoginRequest.Secret,
		Identifier: userLoginRequest.Identifier,
		Type:       string(userLoginRequest.Type),
	}
	userLoginRespDto, err := ctrl.AuthService.LoginUser(ctx, userLoginDto)
	if err != nil {
		return controller_uitl.ExposeError(response, err,
			auth_enum.CODE.UNKNOWN_CREDENTIAL,
			auth_enum.CODE.PASSWORD_WRONG,
			auth_enum.CODE.EMAIL_CODE_WRONG,
			auth_enum.CODE.GOOGLE_2FA_WRONG,
			auth_enum.CODE.OAUTH_WRONG,
		)
	}
	// TODO expire config
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
		return controller_uitl.ExposeError(response, err,
			auth_enum.CODE.UNKNOWN_SESSION,
		)
	}
	return response.Success()
}
