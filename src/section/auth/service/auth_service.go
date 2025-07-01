package service

import (
	"MVC_DI/section/auth/dto"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=service.go -destination=..\..\..\mock\auth\service\service_mock.go -package=service_mock
type AuthService interface {
	// LoginUser
	//
	// error list: enum.CODE_GRPC_ERROR, enum.CODE_UNKNOWN_CREDENTIAL, enum.CODE_PASSWORD_WRONG, enum.CODE_EMAIL_CODE_WRONG, enum.CODE_GOOGLE_2FA_WRONG, enum.CODE_OAUTH_WRONG
	LoginUser(ctx *gin.Context, userLoginDto dto.UserLoginDto) (*dto.UserLoginRespDto, error)
	// LogoutUser
	//
	// error list: enum.CODE_GRPC_ERROR, enum.CODE_UNKNOWN_SESSION
	LogoutUser(ctx *gin.Context, sessionId int64) error
	// DEFINE METHODS
}
