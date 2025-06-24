package auth_mapper_impl

import (
	"MVC_DI/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	auth_mapper "MVC_DI/section/auth/mapper"
	"context"
	"errors"

	"google.golang.org/grpc"
)

type AuthMapperImpl struct {
	conn *grpc.ClientConn
}

// InvalidSession implements auth_mapper.AuthMapper.
func (a *AuthMapperImpl) InvalidSession(ctx context.Context, id int64) error {
	client := proto.NewAuthSessionServiceClient(a.conn)
	request := &proto.InvalidateSessionRequest{SessionId: id}
	_, err := client.InvalidateSession(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthMapperImpl) CreateSession(ctx context.Context, dto auth_dto.CreateSessionDto) (*int64, error) {
	client := proto.NewAuthSessionServiceClient(a.conn)
	request := &proto.CreateAuthSessionRequest{
		UserId: dto.UserId,
	}
	response, err := client.CreateAuthSession(ctx, request)
	if err != nil {
		return nil, err
	}
	sessionID := response.GetSessionId()
	return &sessionID, nil
}

func (a AuthMapperImpl) GetCredentialsByIdentifierAndType(ctx context.Context, dto auth_dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error) {
	client := proto.NewAuthCredentialServiceClient(a.conn)
	val := proto.CredentialType_value[dto.Type]
	credentialType := proto.CredentialType(val)
	request := proto.GetAuthCredentialsRequest{Identifier: &dto.Identifier, Type: &credentialType}
	response, err := client.GetAuthCredentials(ctx, &request)
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
