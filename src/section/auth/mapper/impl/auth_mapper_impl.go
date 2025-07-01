package impl

import (
	"MVC_DI/gen/proto"
	global_enum "MVC_DI/global/enum"
	"MVC_DI/global/model"
	"MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	"MVC_DI/section/auth/mapper"
	"context"
)

type AuthMapperImpl struct {
	AuthSessionServiceClient    proto.AuthSessionServiceClient
	AuthCredentialServiceClient proto.AuthCredentialServiceClient
}

func (a AuthMapperImpl) CreateSession(ctx context.Context, dto dto.CreateSessionDto) (*int64, error) {
	request := &proto.CreateAuthSessionRequest{
		UserId: dto.UserId,
	}
	response, err := a.AuthSessionServiceClient.CreateAuthSession(ctx, request)
	if err != nil {
		return nil, model.NewAppError().WithStatusKey(global_enum.GRPC_ERROR{}).WithDetail(err)
	}
	sessionID := response.GetSessionId()
	return &sessionID, nil
}

func (a AuthMapperImpl) GetCredentialsByIdentifierAndType(ctx context.Context, dto dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error) {
	val := proto.CredentialType_value[dto.Type]
	credentialType := proto.CredentialType(val)
	request := proto.GetAuthCredentialsRequest{Identifier: &dto.Identifier, Type: &credentialType}
	response, err := a.AuthCredentialServiceClient.GetAuthCredentials(ctx, &request)
	if err != nil {
		return nil, model.NewAppError().WithStatusKey(global_enum.GRPC_ERROR{}).WithDetail(err)
	}
	if len(response.GetCredentials()) == 0 {
		return nil, model.NewAppError().WithStatusKeyOptionalMap(auth_enum.UNKNOWN_CREDENTIAL{}, &auth_enum.AUTH_STATUS_MAP)
	}
	return response.GetCredentials(), nil
}

// INTERFACE
var _ mapper.AuthMapper = (*AuthMapperImpl)(nil)
