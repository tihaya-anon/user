package auth_service_impl

import (
	"MVC_DI/gen/proto"
	global_model "MVC_DI/global/model"
	avro_serializer "MVC_DI/infra/avro/serializer"
	event_mapper "MVC_DI/infra/event/mapper"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_mapper "MVC_DI/section/auth/mapper"
	auth_service "MVC_DI/section/auth/service"
	"MVC_DI/security/jwt"
	"MVC_DI/security/jwt/claims"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServiceImpl struct {
	AuthMapper     auth_mapper.AuthMapper
	EventMapper    event_mapper.EventMapper
	MatchService   auth_service.MatchService
	AvroSerializer avro_serializer.IAvroSerializer
	Logger         *logrus.Logger
}

// LogoutUser implements auth_service.AuthService.
//
// error list: enum.CODE.GRPC_ERROR, auth_enum.CODE.UNKNOWN_SESSION
func (a *AuthServiceImpl) LogoutUser(ctx *gin.Context, sessionId int64) error {
	// TODO: dynamically decide the trigger mode
	addLoginAuditLogRequest := &proto.InvalidateSessionRequest{SessionId: sessionId}
	subject, schemaId, payload, err := a.AvroSerializer.SerializeProtoMessage(addLoginAuditLogRequest)
	if err != nil {
		return err
	}
	envelope := &proto.KafkaEnvelope{
		SchemaSubject:        subject,
		SchemaId:             schemaId,
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
// error list: enum.CODE.GRPC_ERROR, auth_enum.CODE.UNKNOWN_CREDENTIAL, auth_enum.CODE.PASSWORD_WRONG, auth_enum.CODE.EMAIL_CODE_WRONG, auth_enum.CODE_2FA_WRONG, auth_enum.CODE.OAUTH_WRONG
func (a *AuthServiceImpl) LoginUser(ctx *gin.Context, userLoginDto auth_dto.UserLoginDto) (*auth_dto.UserLoginRespDto, error) {
	getCredentialsByIdentifierAndTypeDto := auth_dto.GetCredentialsByIdentifierAndTypeDto{Identifier: userLoginDto.Identifier, Type: userLoginDto.Type}
	authCredentials, err := a.AuthMapper.GetCredentialsByIdentifierAndType(ctx, getCredentialsByIdentifierAndTypeDto)
	if err != nil {
		return nil, err
	}

	appErr := global_model.NewAppError()

	authCredential := getActiveCredential(authCredentials)
	if authCredential == nil {
		return nil, appErr.WithCode(auth_enum.CODE.UNKNOWN_CREDENTIAL).WithMessage(auth_enum.MSG.UNKNOWN_CREDENTIAL)
	}

	addLoginAuditLogRequest := &proto.AddLoginAuditLogRequest{
		UserId:     authCredential.GetUserId(),
		LoginTime:  timestamppb.Now(),
		IpAddress:  ctx.Request.RemoteAddr,
		DeviceInfo: "info",
	}
	matched := a.matchCredential(userLoginDto, authCredential, appErr, addLoginAuditLogRequest)

	if !matched {
		a.submitAuditLogEvent(ctx, addLoginAuditLogRequest)
		return nil, appErr
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

	response := &auth_dto.UserLoginRespDto{
		SessionId: *sessionId,
		Token:     token,
	}
	addLoginAuditLogRequest.Result = proto.LoginResult_SUCCESS
	a.submitAuditLogEvent(ctx, addLoginAuditLogRequest)
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

func (a *AuthServiceImpl) matchCredential(userLoginDto auth_dto.UserLoginDto, authCredential *proto.AuthCredential, appErr *global_model.AppError, addLoginAuditLogRequest *proto.AddLoginAuditLogRequest) bool {
	switch authCredential.Type {
	case proto.CredentialType_PASSWORD:
		appErr.WithCode(auth_enum.CODE.PASSWORD_WRONG).WithMessage(auth_enum.MSG.PASSWORD_WRONG)
		addLoginAuditLogRequest.Result = proto.LoginResult_FAIL_PASSWORD
		return a.MatchService.MatchPassword(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
	case proto.CredentialType_EMAIL_CODE:
		appErr.WithCode(auth_enum.CODE.EMAIL_CODE_WRONG).WithMessage(auth_enum.MSG.EMAIL_CODE_WRONG)
		addLoginAuditLogRequest.Result = proto.LoginResult_FAIL_EMAIL_CODE
		return a.MatchService.MatchEmailCode(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
	case proto.CredentialType__2FA:
		appErr.WithCode(auth_enum.CODE.GOOGLE_2FA_WRONG).WithMessage(auth_enum.MSG.GOOGLE_2FA_WRONG)
		addLoginAuditLogRequest.Result = proto.LoginResult_FAIL_2FA
		return a.MatchService.MatchGoogle2FA(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
	case proto.CredentialType_OAUTH:
		appErr.WithCode(auth_enum.CODE.OAUTH_WRONG).WithMessage(auth_enum.MSG.OAUTH_WRONG)
		addLoginAuditLogRequest.Result = proto.LoginResult_FAIL_OAUTH
		return a.MatchService.MatchOauth(userLoginDto.Identifier, userLoginDto.Secret, authCredential.GetSecret())
	default:
		appErr.WithCode(auth_enum.CODE.UNKNOWN_CREDENTIAL).WithMessage(auth_enum.MSG.UNKNOWN_CREDENTIAL)
		addLoginAuditLogRequest.Result = proto.LoginResult_LOGIN_RESULT_UNSPECIFIED
		return false
	}
}

func (a *AuthServiceImpl) submitAuditLogEvent(ctx *gin.Context, request *proto.AddLoginAuditLogRequest) {
	subject, schemaId, payload, err := a.AvroSerializer.SerializeProtoMessage(request)
	if err != nil {
		return
	}
	envelope := &proto.KafkaEnvelope{
		SchemaSubject:        subject,
		SchemaId:             schemaId,
		Priority:             proto.Priority_LOW,
		Payload:              payload,
		DeliveryMode:         proto.DeliveryMode_PULL,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
	}
	err = a.EventMapper.SubmitEvent(ctx, envelope)
	if err != nil {
		// log error
	}
}

// INTERFACE
var _ auth_service.AuthService = (*AuthServiceImpl)(nil)
