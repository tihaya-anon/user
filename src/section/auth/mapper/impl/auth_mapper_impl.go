package auth_mapper_impl

import (
	"MVC_DI/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	auth_mapper "MVC_DI/section/auth/mapper"
	"context"
	"errors"
	"fmt"
)

type AuthMapperImpl struct {
	KafkaEventServiceClient     proto.KafkaEventServiceClient
	AuthSessionServiceClient    proto.AuthSessionServiceClient
	AuthCredentialServiceClient proto.AuthCredentialServiceClient
}

// InvalidSession implements auth_mapper.AuthMapper.
func (a *AuthMapperImpl) InvalidSession(ctx context.Context, id int64) error {
	// TODO dynamically get the trigger mode
	envelope := proto.KafkaEnvelope{DeliveryMode: proto.DeliveryMode_PUSH, TriggerMode: proto.TriggerMode_ASYNC}
	response, err := a.KafkaEventServiceClient.SubmitEvent(ctx, &proto.SubmitEventRequest{Envelope: &envelope})
	if err != nil {
		return err
	}

	switch envelope.GetTriggerMode() {
	case proto.TriggerMode_ASYNC:
		return nil

	case proto.TriggerMode_SYNC:
		if response.GetStatus() != proto.EventStatus_PROCESSED_SUCCESS {
			return fmt.Errorf("sync event failed: %s", response.GetStatus().String())
		}
		return nil

	default:
		return fmt.Errorf("invalid trigger mode: %s", envelope.GetTriggerMode().String())
	}
}

func (a AuthMapperImpl) CreateSession(ctx context.Context, dto auth_dto.CreateSessionDto) (*int64, error) {
	request := &proto.CreateAuthSessionRequest{
		UserId: dto.UserId,
	}
	response, err := a.AuthSessionServiceClient.CreateAuthSession(ctx, request)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if len(response.GetCredentials()) == 0 {
		return nil, errors.New("credential not found")
	}
	return response.GetCredentials(), nil
}

// INTERFACE
var _ auth_mapper.AuthMapper = (*AuthMapperImpl)(nil)
