package auth_service_impl

import (
	"MVC_DI/gen/proto"
	event_mapper "MVC_DI/global/infra/event/mapper"
	"MVC_DI/global/infra/schema"
	global_model "MVC_DI/global/model"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/security/claims"
	"MVC_DI/security/jwt"
	payload_util "MVC_DI/util/payload"
	"context"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	AuthMapper   auth_mapper.AuthMapper
	EventMapper  event_mapper.EventMapper
	MatchService auth_service.MatchService
}

// LogoutUser implements auth_service.AuthService.
//
// error list: enum.CODE.GRPC_ERROR, auth_enum.CODE.UNKNOWN_SESSION
func (a *AuthServiceImpl) LogoutUser(ctx *gin.Context, sessionId int64) error {
	// TODO: dynamically decide the trigger mode
	request := &proto.InvalidateSessionRequest{SessionId: sessionId}
	native, err := payload_util.ProtoToNative(request)
	if err != nil {
		return err
	}
	codec, _, err := schema.SchemaManager.GetOrLoadCodecByObject(request)
	if err != nil {
		return err
	}
	payload, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		return err
	}
	envelope := &proto.KafkaEnvelope{
		Priority:             proto.Priority_HIGH,
		Payload:              payload,
		DeliveryMode:         proto.DeliveryMode_PUSH,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
	}
	err = a.EventMapper.SubmitEvent(ctx, envelope)
	return err
}

// LoginUser implements auth_service.AuthService.
//
// error list: enum.CODE.GRPC_ERROR, auth_enum.CODE.UNKNOWN_CREDENTIAL, auth_enum.CODE.PASSWORD_WRONG, auth_enum.CODE.EMAIL_CODE_WRONG, auth_enum.CODE.GOOGLE_2FA_WRONG, auth_enum.CODE.OAUTH_WRONG
func (a AuthServiceImpl) LoginUser(ctx context.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error) {
	getCredentialsByIdentifierAndTypeDto := auth_dto.GetCredentialsByIdentifierAndTypeDto{Identifier: userLoginDto.Identifier, Type: userLoginDto.Type}
	authCredentials, err := a.AuthMapper.GetCredentialsByIdentifierAndType(ctx, getCredentialsByIdentifierAndTypeDto)
	if err != nil {
		return nil, err
	}
	matched := false
	appErr := global_model.NewAppError()
	authCredential := &proto.AuthCredential{}
	authCredential = nil
	for _, credential := range authCredentials {
		if !credential.GetIsActive() {
			continue
		}
		authCredential = credential
		break
	}
	if authCredential == nil {
		return nil, appErr.WithCode(auth_enum.CODE.UNKNOWN_CREDENTIAL).WithMessage(auth_enum.MSG.UNKNOWN_CREDENTIAL)
	}
	switch authCredential.Type {
	case proto.CredentialType_PASSWORD:
		matched = a.MatchService.MatchPassword(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
		appErr.WithCode(auth_enum.CODE.PASSWORD_WRONG).WithMessage(auth_enum.MSG.PASSWORD_WRONG)
	case proto.CredentialType_EMAIL_CODE:
		matched = a.MatchService.MatchEmailCode(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
		appErr.WithCode(auth_enum.CODE.EMAIL_CODE_WRONG).WithMessage(auth_enum.MSG.EMAIL_CODE_WRONG)
	case proto.CredentialType_GOOGLE_2FA:
		matched = a.MatchService.MatchGoogle2FA(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
		appErr.WithCode(auth_enum.CODE.GOOGLE_2FA_WRONG).WithMessage(auth_enum.MSG.GOOGLE_2FA_WRONG)
	case proto.CredentialType_OAUTH:
		matched = a.MatchService.MatchOauth(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
		appErr.WithCode(auth_enum.CODE.OAUTH_WRONG).WithMessage(auth_enum.MSG.OAUTH_WRONG)
	}
	if !matched {
		return nil, appErr
	}
	createSessionDto := auth_dto.CreateSessionDto{UserId: authCredential.GetUserId()}
	sessionId, err := a.AuthMapper.CreateSession(ctx, createSessionDto)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateJWT(claims.UserClaim{
		UserId: authCredential.GetUserId(),
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
