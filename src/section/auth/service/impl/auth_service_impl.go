package impl

import (
	"MVC_DI/async"
	"MVC_DI/gen/proto"
	"MVC_DI/global/model"
	"MVC_DI/section/auth/dto"
	"MVC_DI/section/auth/enum"
	"MVC_DI/section/auth/event"
	"MVC_DI/section/auth/mapper"
	"MVC_DI/section/auth/service"
	"MVC_DI/security/jwt"
	"MVC_DI/security/jwt/claims"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthServiceImpl struct {
	AuthMapper         mapper.AuthMapper
	VerifyService      service.VerifyService
	AuthEventPublisher event.AuthEventPublisher
	Logger             *logrus.Logger
}

// LogoutUser implements service.AuthService.
//
// error list: enum.CODE_GRPC_ERROR, enum.CODE_UNKNOWN_SESSION
func (a *AuthServiceImpl) LogoutUser(ctx *gin.Context, sessionId int64) error {
	return a.AuthEventPublisher.PublishInvalidSession(ctx, sessionId)
}

// LoginUser implements service.AuthService.
//
// error list: enum.CODE_GRPC_ERROR, enum.CODE_UNKNOWN_CREDENTIAL, enum.CODE_PASSWORD_WRONG, enum.CODE_EMAIL_CODE_WRONG, enum.CODE_2FA_WRONG, enum.CODE_OAUTH_WRONG
func (a *AuthServiceImpl) LoginUser(ctx *gin.Context, userLoginDto dto.UserLoginDto) (*dto.UserLoginRespDto, error) {
	getCredentialsByIdentifierAndTypeDto := dto.GetCredentialsByIdentifierAndTypeDto{Identifier: userLoginDto.Identifier, Type: userLoginDto.Type}
	authCredentials, err := a.AuthMapper.GetCredentialsByIdentifierAndType(ctx, getCredentialsByIdentifierAndTypeDto)
	if err != nil {
		return nil, err
	}

	authCredential := getActiveCredential(authCredentials)
	if authCredential == nil {
		return nil, model.NewAppError().WithStatusKey(enum.UNKNOWN_CREDENTIAL{})
	}
	ok, result, err := a.VerifyService.Verify(userLoginDto, authCredential)

	if !ok {
		cloneCtx := context.Background()
		ipAddress := ctx.Request.RemoteAddr
		async.AsyncCtx(cloneCtx, a.Logger, func(c context.Context) {
			publishLoginAuditDto := &dto.PublishLoginAuditDto{
				IpAddress:  ipAddress,
				UserId:     authCredential.GetUserId(),
				DeviceInfo: "deviceInfo",
				Result:     result,
			}
			_ = a.AuthEventPublisher.PublishLoginAudit(c, publishLoginAuditDto)
		})
		return nil, err
	}

	token, err := jwt.GenerateJWT(claims.UserClaim{
		UserId: authCredential.GetUserId(),
	})
	if err != nil {
		return nil, err
	}

	createSessionDto := dto.CreateSessionDto{UserId: authCredential.GetUserId()}
	sessionId, err := a.AuthMapper.CreateSession(ctx, createSessionDto)
	if err != nil {
		return nil, err
	}
	cloneCtx := context.Background()
	ipAddress := ctx.Request.RemoteAddr
	async.AsyncCtx(cloneCtx, a.Logger, func(c context.Context) {
		publishLoginAuditDto := &dto.PublishLoginAuditDto{
			IpAddress:  ipAddress,
			UserId:     authCredential.GetUserId(),
			DeviceInfo: "deviceInfo",
			Result:     proto.LoginResult_SUCCESS,
		}
		_ = a.AuthEventPublisher.PublishLoginAudit(c, publishLoginAuditDto)
	})
	response := &dto.UserLoginRespDto{
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
var _ service.AuthService = (*AuthServiceImpl)(nil)
