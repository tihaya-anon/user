package auth_service

import (
	auth_dto "MVC_DI/section/auth/dto"
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=auth_service.go -destination=..\..\..\mock\auth\service\auth_service_mock.go -package=auth_service_mock
type AuthService interface {
	LoginUser(ctx context.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error)
	LogoutUser(ctx *gin.Context, sessionId int64) error
	// DEFINE METHODS
}
