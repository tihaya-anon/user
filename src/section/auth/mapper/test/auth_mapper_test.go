package mapper_test

//go:generate mockgen -destination=../../../../mock/gen/proto/auth_session_mock.go -package=proto_mock MVC_DI/gen/proto AuthSessionServiceClient
//go:generate mockgen -destination=../../../../mock/gen/proto/auth_credential_mock.go -package=proto_mock MVC_DI/gen/proto AuthCredentialServiceClient

import (
	"context"
	"errors"
	"testing"

	"MVC_DI/gen/proto"
	global_enum "MVC_DI/global/enum"
	"MVC_DI/global/model"
	proto_mock "MVC_DI/mock/gen/proto"
	"MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	"MVC_DI/section/auth/mapper/impl"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateSession_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := proto_mock.NewMockAuthSessionServiceClient(ctrl)
	authMapper := &impl.AuthMapperImpl{
		AuthSessionServiceClient: mockSession,
	}

	dto := dto.CreateSessionDto{UserId: 42}
	ctx := context.Background()

	sessionID := int64(1001)
	mockSession.EXPECT().
		CreateAuthSession(ctx, gomock.Any()).
		Return(&proto.CreateAuthSessionResponse{SessionId: sessionID}, nil)

	result, err := authMapper.CreateSession(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, &sessionID, result)
}

func Test_CreateSession_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := proto_mock.NewMockAuthSessionServiceClient(ctrl)
	authMapper := &impl.AuthMapperImpl{
		AuthSessionServiceClient: mockSession,
	}

	dto := dto.CreateSessionDto{UserId: 42}
	ctx := context.Background()
	errMsg := "create session failed"
	mockSession.EXPECT().
		CreateAuthSession(ctx, gomock.Any()).
		Return(nil, errors.New(errMsg))

	result, err := authMapper.CreateSession(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, model.NewAppError().WithStatusKey(global_enum.GRPC_ERROR{}).WithDetail(errors.New(errMsg)).Error())
}

func Test_GetCredentialsByIdentifierAndType_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCredential := proto_mock.NewMockAuthCredentialServiceClient(ctrl)
	authMapper := &impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := dto.GetCredentialsByIdentifierAndTypeDto{
		Identifier: "email@example.com",
		Type:       "PASSWORD",
	}
	ctx := context.Background()

	cred := &proto.AuthCredential{UserId: 1, IsActive: true}
	mockCredential.EXPECT().
		GetAuthCredentials(ctx, gomock.Any()).
		Return(&proto.GetAuthCredentialsResponse{Credentials: []*proto.AuthCredential{cred}}, nil)

	result, err := authMapper.GetCredentialsByIdentifierAndType(ctx, dto)
	assert.NoError(t, err)
	assert.Equal(t, []*proto.AuthCredential{cred}, result)
}

func Test_GetCredentialsByIdentifierAndType_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCredential := proto_mock.NewMockAuthCredentialServiceClient(ctrl)
	authMapper := &impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := dto.GetCredentialsByIdentifierAndTypeDto{
		Identifier: "email@example.com",
		Type:       "PASSWORD",
	}
	ctx := context.Background()

	mockCredential.EXPECT().
		GetAuthCredentials(ctx, gomock.Any()).
		Return(&proto.GetAuthCredentialsResponse{Credentials: []*proto.AuthCredential{}}, nil)

	result, err := authMapper.GetCredentialsByIdentifierAndType(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, model.NewAppError().WithStatusKeyOptionalMap(auth_enum.UNKNOWN_CREDENTIAL{}, &auth_enum.AUTH_STATUS_MAP).Error())
}

func Test_GetCredentialsByIdentifierAndType_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCredential := proto_mock.NewMockAuthCredentialServiceClient(ctrl)
	authMapper := &impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := dto.GetCredentialsByIdentifierAndTypeDto{
		Identifier: "email@example.com",
		Type:       "PASSWORD",
	}
	ctx := context.Background()
	errMsg := "rpc failed"
	mockCredential.EXPECT().
		GetAuthCredentials(ctx, gomock.Any()).
		Return(nil, errors.New(errMsg))

	result, err := authMapper.GetCredentialsByIdentifierAndType(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, model.NewAppError().WithStatusKey(global_enum.GRPC_ERROR{}).WithDetail(errors.New(errMsg)).Error())
}
