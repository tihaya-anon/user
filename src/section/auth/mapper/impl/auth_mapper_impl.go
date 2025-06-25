package auth_mapper_impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/enum"
	global_model "MVC_DI/global/model"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_mapper "MVC_DI/section/auth/mapper"
	"context"
)

type AuthMapperImpl struct {
	KafkaEventServiceClient     proto.KafkaEventServiceClient
	AuthSessionServiceClient    proto.AuthSessionServiceClient
	AuthCredentialServiceClient proto.AuthCredentialServiceClient
}

// InvalidSession implements auth_mapper.AuthMapper.
func (a *AuthMapperImpl) InvalidSession(ctx context.Context, envelope *proto.KafkaEnvelope) error {
	response, err := a.KafkaEventServiceClient.SubmitEvent(ctx, &proto.SubmitEventRequest{Envelope: envelope})
	if err != nil {
		return err
	}

	switch envelope.GetTriggerModeRequested() {
	case proto.TriggerMode_ASYNC:
		return nil

	case proto.TriggerMode_SYNC:
		if response.GetStatus() != proto.EventStatus_PROCESSED_SUCCESS {
			return global_model.NewAppError().WithCode(enum.CODE.GRPC_ERROR).WithDetail(response.GetStatus().String())
		}
		return nil

	default:
		return global_model.NewAppError().WithCode(auth_enum.CODE.UNKNOWN_TRIGGER_MODE).WithMessage(auth_enum.MSG.UNKNOWN_TRIGGER_MODE).WithDetail(envelope.GetTriggerModeRequested().String())
	}
}

func (a AuthMapperImpl) CreateSession(ctx context.Context, dto auth_dto.CreateSessionDto) (*int64, error) {
	request := &proto.CreateAuthSessionRequest{
		UserId: dto.UserId,
	}
	response, err := a.AuthSessionServiceClient.CreateAuthSession(ctx, request)
	if err != nil {
		return nil, global_model.NewAppError().WithCode(enum.CODE.GRPC_ERROR).WithMessage(err.Error())
	}
	sessionID := response.GetSessionId()
	return &sessionID, nil
}

func (a AuthMapperImpl) GetCredentialsByIdentifierAndType(ctx context.Context, dto auth_dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error) {
	val := proto.CredentialType_value[dto.Type]
	credentialType := proto.CredentialType(val)
	request := proto.GetAuthCredentialsRequest{Identifier: &dto.Identifier, Type: &credentialType}
	response, err := a.AuthCredentialServiceClient.GetAuthCredentials(ctx, &request)
	if err != nil {
		return nil, global_model.NewAppError().WithCode(enum.CODE.GRPC_ERROR).WithMessage(err.Error())
	}
	if len(response.GetCredentials()) == 0 {
		return nil, global_model.NewAppError().WithCode(auth_enum.CODE.UNKNOWN_CREDENTIAL).WithMessage(auth_enum.MSG.UNKNOWN_CREDENTIAL)
	}
	return response.GetCredentials(), nil
}

// INTERFACE
var _ auth_mapper.AuthMapper = (*AuthMapperImpl)(nil)
