package auth_mapper

import (
	"MVC_DI/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	"context"
)

//go:generate mockgen -source=auth_mapper.go -destination=..\..\..\mock\auth\mapper\auth_mapper_mock.go -package=auth_mapper_mock
type AuthMapper interface {
	// GetCredentialsByIdentifierAndType
	// 
	// error list:
	// enum.CODE.GRPC_ERROR, auth_enum.CODE.UNKNOWN_CREDENTIAL
	GetCredentialsByIdentifierAndType(ctx context.Context, dto auth_dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error)
	// InvalidSession
	// 
	// error list:
	// enum.CODE.GRPC_ERROR
	CreateSession(ctx context.Context, dto auth_dto.CreateSessionDto) (*int64, error)
	// DEFINE METHODS
}
