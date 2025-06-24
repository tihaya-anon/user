package auth_mapper

import (
	"MVC_DI/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	"context"
)

//go:generate mockgen -source=auth_mapper.go -destination=..\..\..\mock\auth\mapper\auth_mapper_mock.go -package=auth_mapper_mock
type AuthMapper interface {
	GetCredentialsByIdentifierAndType(ctx context.Context, dto auth_dto.GetCredentialsByIdentifierAndTypeDto) ([]*proto.AuthCredential, error)
	CreateSession(ctx context.Context, dto auth_dto.CreateSessionDto) (*int64, error)
	InvalidSession(ctx context.Context, sessionId int64) error
	// DEFINE METHODS
}
