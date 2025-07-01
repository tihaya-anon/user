package auth_service

import (
	auth_dto "MVC_DI/section/auth/dto"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=auth_service.go -destination=..\..\..\mock\auth\service\auth_service_mock.go -package=auth_service_mock
type AuthService interface {
	// LoginUser
	//
	// error list: enum.CODE_GRPC_ERROR, auth_enum.CODE_UNKNOWN_CREDENTIAL, auth_enum.CODE_PASSWORD_WRONG, auth_enum.CODE_EMAIL_CODE_WRONG, auth_enum.CODE_GOOGLE_2FA_WRONG, auth_enum.CODE_OAUTH_WRONG
	LoginUser(ctx *gin.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error)
	// LogoutUser
	//
	// error list: enum.CODE_GRPC_ERROR, auth_enum.CODE_UNKNOWN_SESSION
	LogoutUser(ctx *gin.Context, sessionId int64) error
	// DEFINE METHODS
}
