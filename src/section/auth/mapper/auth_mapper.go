package mapper

import (
	"MVC_DI/gen/proto"
	"MVC_DI/section/auth/dto"
	"context"
)

//go:generate mockgen -source=mapper.go -destination=..\..\..\mock\auth\mapper\mapper_mock.go -package=mapper_mock
type AuthMapper interface {
	// GetCredentialsByIdentifierAndType
	//
	// error list:
	// enum.CODE_GRPC_ERROR, auth_enum.CODE_UNKNOWN_CREDENTIAL
	GetCredentialsByIdentifierAndType(ctx context.Context, dto dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error)
	// InvalidSession
	//
	// error list:
	// enum.CODE_GRPC_ERROR
	CreateSession(ctx context.Context, dto dto.CreateSessionDto) (*int64, error)
	// DEFINE METHODS
}
