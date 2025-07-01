package auth_service_impl

import (
	"MVC_DI/async"
	"MVC_DI/gen/proto"
	global_model "MVC_DI/global/model"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_event_publisher "MVC_DI/section/auth/event"
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/security/jwt"
	"MVC_DI/security/jwt/claims"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthServiceImpl struct {
	AuthMapper         auth_mapper.AuthMapper
	VerifyService      auth_service.VerifyService
	AuthEventPublisher auth_event_publisher.AuthEventPublisher
	Logger             *logrus.Logger
}

// LogoutUser implements auth_service.AuthService.
//
// error list: enum.CODE_GRPC_ERROR, auth_enum.CODE_UNKNOWN_SESSION
func (a *AuthServiceImpl) LogoutUser(ctx *gin.Context, sessionId int64) error {
	return a.AuthEventPublisher.PublishInvalidSession(ctx, sessionId)
}

// LoginUser implements auth_service.AuthService.
//
// error list: enum.CODE_GRPC_ERROR, auth_enum.CODE_UNKNOWN_CREDENTIAL, auth_enum.CODE_PASSWORD_WRONG, auth_enum.CODE_EMAIL_CODE_WRONG, auth_enum.CODE_2FA_WRONG, auth_enum.CODE_OAUTH_WRONG
func (a *AuthServiceImpl) LoginUser(ctx *gin.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error) {
	getCredentialsByIdentifierAndTypeDto := auth_dto.GetCredentialsByIdentifierAndTypeDto{Identifier: userLoginDto.Identifier, Type: userLoginDto.Type}
	authCredentials, err := a.AuthMapper.GetCredentialsByIdentifierAndType(ctx, getCredentialsByIdentifierAndTypeDto)
	if err != nil {
		return nil, err
	}

	authCredential := getActiveCredential(authCredentials)
	if authCredential == nil {
		return nil, global_model.NewAppError().WithStatusKey(auth_enum.UNKNOWN_CREDENTIAL{})
	}
	ok, result, err := a.VerifyService.Verify(userLoginDto, authCredential)

	if !ok {
		cloneCtx := context.Background()
		ipAddress := ctx.Request.RemoteAddr
		async.AsyncCtx(cloneCtx, a.Logger, func(c context.Context) {
			publishLoginAuditDto := &auth_dto.PublishLoginAuditDto{
				IpAddress:  ipAddress,
				UserId:     authCredential.GetUserId(),
				DeviceInfo: "deviceInfo",
				Result:     result,
			}
			_ = a.AuthEventPublisher.PublishLoginAudit(ctx, publishLoginAuditDto)
		})
		return nil, err
	}

	token, err := jwt.GenerateJWT(claims.UserClaim{
		UserId: authCredential.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	createSessionDto := auth_dto.CreateSessionDto{UserId: authCredential.GetUserId()}
	sessionId, err := a.AuthMapper.CreateSession(ctx, createSessionDto)
	if err != nil {
		return nil, err
	}
	cloneCtx := context.Background()
	ipAddress := ctx.Request.RemoteAddr
	async.AsyncCtx(cloneCtx, a.Logger, func(c context.Context) {
		publishLoginAuditDto := &auth_dto.PublishLoginAuditDto{
			IpAddress:  ipAddress,
			UserId:     authCredential.GetUserId(),
			DeviceInfo: "deviceInfo",
			Result:     proto.LoginResult_SUCCESS,
		}
		_ = a.AuthEventPublisher.PublishLoginAudit(ctx, publishLoginAuditDto)
	})
	response := &auth_dto.UserLoginRespDto{
		SessionId: *sessionId,
		Token:     token,
	}
	return response, nil
}

func getActiveCredential(credentials []*proto.AuthCredential) *proto.AuthCredential {
	for _, credential := range credentials {
		if credential.IsActive {
			return credential
		}
	}
	return nil
}

// INTERFACE
var _ auth_service.AuthService = (*AuthServiceImpl)(nil)
