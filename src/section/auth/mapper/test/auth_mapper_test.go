package auth_mapper_test

import (
	"context"
	"errors"
	"testing"

	"MVC_DI/gen/proto"
	proto_mock "MVC_DI/mock/gen/proto"
	auth_dto "MVC_DI/section/auth/dto"
	auth_mapper_impl "MVC_DI/section/auth/mapper/impl"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_InvalidSession_AsyncMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKafka := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		KafkaEventServiceClient: mockKafka,
	}

	ctx := context.Background()

	mockKafka.EXPECT().
		SubmitEvent(ctx, gomock.Any()).
		Return(&proto.SubmitEventResponse{Status: proto.EventStatus_STATUS_UNSPECIFIED}, nil)

	err := authMapper.InvalidSession(ctx, 123)
	assert.NoError(t, err)
}

func Test_InvalidSession_SyncMode_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKafka := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		KafkaEventServiceClient: mockKafka,
	}

	ctx := context.Background()

	// simulate sync mode by altering envelope manually via mock expectation
	mockKafka.EXPECT().
		SubmitEvent(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, req *proto.SubmitEventRequest, _ ...any) (*proto.SubmitEventResponse, error) {
			req.Envelope.TriggerMode = proto.TriggerMode_SYNC
			return &proto.SubmitEventResponse{Status: proto.EventStatus_PROCESSED_SUCCESS}, nil
		})

	err := authMapper.InvalidSession(ctx, 123)
	assert.NoError(t, err)
}

func Test_InvalidSession_SyncMode_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockKafka := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		KafkaEventServiceClient: mockKafka,
	}

	ctx := context.Background()

	mockKafka.EXPECT().
		SubmitEvent(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, req *proto.SubmitEventRequest, _ ...any) (*proto.SubmitEventResponse, error) {
			req.Envelope.TriggerMode = proto.TriggerMode_SYNC
			return &proto.SubmitEventResponse{Status: proto.EventStatus_PROCESSED_FAILED}, nil
		})

	err := authMapper.InvalidSession(ctx, 123)
	assert.EqualError(t, err, "sync event failed: PROCESSED_FAILED")
}

func Test_CreateSession_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSession := proto_mock.NewMockAuthSessionServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		AuthSessionServiceClient: mockSession,
	}

	dto := auth_dto.CreateSessionDto{UserId: 42}
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
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		AuthSessionServiceClient: mockSession,
	}

	dto := auth_dto.CreateSessionDto{UserId: 42}
	ctx := context.Background()

	mockSession.EXPECT().
		CreateAuthSession(ctx, gomock.Any()).
		Return(nil, errors.New("create session failed"))

	result, err := authMapper.CreateSession(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, "create session failed")
}

func Test_GetCredentialsByIdentifierAndType_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCredential := proto_mock.NewMockAuthCredentialServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := auth_dto.GetCredentialsByIdentifierAndTypeDto{
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
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := auth_dto.GetCredentialsByIdentifierAndTypeDto{
		Identifier: "email@example.com",
		Type:       "PASSWORD",
	}
	ctx := context.Background()

	mockCredential.EXPECT().
		GetAuthCredentials(ctx, gomock.Any()).
		Return(&proto.GetAuthCredentialsResponse{Credentials: []*proto.AuthCredential{}}, nil)

	result, err := authMapper.GetCredentialsByIdentifierAndType(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, "credential not found")
}

func Test_GetCredentialsByIdentifierAndType_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCredential := proto_mock.NewMockAuthCredentialServiceClient(ctrl)
	authMapper := &auth_mapper_impl.AuthMapperImpl{
		AuthCredentialServiceClient: mockCredential,
	}

	dto := auth_dto.GetCredentialsByIdentifierAndTypeDto{
		Identifier: "email@example.com",
		Type:       "PASSWORD",
	}
	ctx := context.Background()

	mockCredential.EXPECT().
		GetAuthCredentials(ctx, gomock.Any()).
		Return(nil, errors.New("rpc failed"))

	result, err := authMapper.GetCredentialsByIdentifierAndType(ctx, dto)
	assert.Nil(t, result)
	assert.EqualError(t, err, "rpc failed")
}
