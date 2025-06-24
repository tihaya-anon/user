package auth_service_impl

import (
	"MVC_DI/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/security"
	"MVC_DI/security/claims"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	AuthMapper   auth_mapper.AuthMapper
	MatchService auth_service.MatchService
}

// LogoutUser implements auth_service.AuthService.
func (a *AuthServiceImpl) LogoutUser(ctx *gin.Context, sessionId int64) error {
	err := a.AuthMapper.InvalidSession(ctx, sessionId)
	return err
}

// LoginUser implements auth_service.AuthService.
func (a AuthServiceImpl) LoginUser(ctx context.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error) {
	getCredentialsByIdentifierAndTypeDto := auth_dto.GetCredentialsByIdentifierAndTypeDto{Identifier: userLoginDto.Identifier, Type: userLoginDto.Type}
	authCredentials, err := a.AuthMapper.GetCredentialsByIdentifierAndType(ctx, getCredentialsByIdentifierAndTypeDto)
	if err != nil {
		return nil, err
	}
	matched := false
	msg := auth_enum.MSG.UNKNOWN_CREDENTIAL
	authCredential := &proto.AuthCredential{}
	authCredential = nil
	for _, credential := range authCredentials {
		if !credential.IsActive {
			continue
		}
		authCredential = credential
		break
	}
	if authCredential == nil {
		return nil, fmt.Errorf(msg)
	}
	switch authCredential.Type {
	case proto.CredentialType_PASSWORD:
		matched = a.MatchService.MatchPassword(userLoginDto.Secret, authCredential.Secret)
		msg = auth_enum.MSG.PASSWORD_WRONG
	case proto.CredentialType_EMAIL_CODE:
		matched = a.MatchService.MatchEmailCode(userLoginDto.Secret, authCredential.Secret)
		msg = auth_enum.MSG.EMAIL_CODE_WRONG
	case proto.CredentialType_GOOGLE_2FA:
		matched = a.MatchService.MatchGoogle2FA(userLoginDto.Secret, authCredential.Secret)
		msg = auth_enum.MSG.GOOGLE_2FA_WRONG
	case proto.CredentialType_OAUTH:
		matched = a.MatchService.MatchOauth(userLoginDto.Secret, authCredential.Secret)
		msg = auth_enum.MSG.OAUTH_WRONG
	}
	if !matched {
		return nil, fmt.Errorf(msg)
	}
	createSessionDto := auth_dto.CreateSessionDto{UserId: authCredential.UserId}
	sessionId, err := a.AuthMapper.CreateSession(ctx, createSessionDto)
	if err != nil {
		return nil, err
	}
	token, err := security.GenerateJWT(claims.UserClaim{
		UserId: authCredential.UserId,
	})
	if err != nil {
		return nil, err
	}
	response := &auth_dto.UserLoginRespDto{
		SessionId: *sessionId,
		Token:     token,
	}
	return response, nil
}

// INTERFACE
var _ auth_service.AuthService = (*AuthServiceImpl)(nil)
