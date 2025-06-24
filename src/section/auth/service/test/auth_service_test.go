package auth_service_test

import (
	"MVC_DI/gen/proto"
	auth_mapper_mock "MVC_DI/mock/auth/mapper"
	auth_service_mock "MVC_DI/mock/auth/service"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_service_impl "MVC_DI/section/auth/service/impl"
	"MVC_DI/security"
	"MVC_DI/security/claims"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Test: mapper 返回错误
func Test_LoginUser_GetCredentialInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMapper := auth_mapper_mock.NewMockAuthMapper(ctrl)
	mockMatch := auth_service_mock.NewMockMatchService(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		AuthMapper:   mockMapper,
		MatchService: mockMatch,
	}

	ctx := context.Background()
	dto := auth_dto.UserLoginDto{Identifier: "user@example.com", Type: proto.CredentialType_PASSWORD.String()}

	mockMapper.EXPECT().
		GetCredentialsByIdentifierAndType(ctx, gomock.Any()).
		Return(nil, errors.New("internal error"))

	resp, err := svc.LoginUser(ctx, dto)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "internal error")
}

// Test: 找不到可用 credential
func Test_LoginUser_UnknownCredentialType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMapper := auth_mapper_mock.NewMockAuthMapper(ctrl)
	mockMatch := auth_service_mock.NewMockMatchService(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		AuthMapper:   mockMapper,
		MatchService: mockMatch,
	}

	ctx := context.Background()
	dto := auth_dto.UserLoginDto{Identifier: "user@example.com", Type: proto.CredentialType_PASSWORD.String()}

	mockMapper.EXPECT().
		GetCredentialsByIdentifierAndType(ctx, gomock.Any()).
		Return([]*proto.AuthCredential{
			{IsActive: false}, // 全部 inactive
		}, nil)

	resp, err := svc.LoginUser(ctx, dto)

	assert.Nil(t, resp)
	assert.EqualError(t, err, auth_enum.MSG.UNKNOWN_CREDENTIAL)
}

// Test: session 创建失败
func Test_LoginUser_CreateSessionInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMapper := auth_mapper_mock.NewMockAuthMapper(ctrl)
	mockMatch := auth_service_mock.NewMockMatchService(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		AuthMapper:   mockMapper,
		MatchService: mockMatch,
	}

	ctx := context.Background()
	dto := auth_dto.UserLoginDto{Identifier: "user@example.com", Secret: "1234", Type: proto.CredentialType_PASSWORD.String()}

	activeCredential := &proto.AuthCredential{
		IsActive: true,
		Secret:   "1234",
		Type:     proto.CredentialType_PASSWORD,
		UserId:   42,
	}

	mockMapper.EXPECT().
		GetCredentialsByIdentifierAndType(ctx, gomock.Any()).
		Return([]*proto.AuthCredential{activeCredential}, nil)

	mockMatch.EXPECT().
		MatchPassword(dto.Identifier, dto.Secret, "1234").
		Return(true)

	mockMapper.EXPECT().
		CreateSession(ctx, gomock.Any()).
		Return(nil, errors.New("db failed"))

	resp, err := svc.LoginUser(ctx, dto)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "db failed")
}

// Test: 登录成功
func Test_LoginUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMapper := auth_mapper_mock.NewMockAuthMapper(ctrl)
	mockMatch := auth_service_mock.NewMockMatchService(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		AuthMapper:   mockMapper,
		MatchService: mockMatch,
	}

	ctx := context.Background()
	dto := auth_dto.UserLoginDto{Identifier: "user@example.com", Secret: "secret", Type: proto.CredentialType_PASSWORD.String()}

	activeCredential := &proto.AuthCredential{
		IsActive: true,
		Secret:   "encoded_secret",
		Type:     proto.CredentialType_PASSWORD,
		UserId:   1001,
	}

	mockMapper.EXPECT().
		GetCredentialsByIdentifierAndType(ctx, gomock.Any()).
		Return([]*proto.AuthCredential{activeCredential}, nil)

	mockMatch.EXPECT().
		MatchPassword(dto.Identifier, dto.Secret, "encoded_secret").
		Return(true)

	sessionId := int64(9876)
	mockMapper.EXPECT().
		CreateSession(ctx, gomock.Any()).
		Return(&sessionId, nil)

	resp, err := svc.LoginUser(ctx, dto)
	assert.NoError(t, err)

	token, err := security.GenerateJWT(claims.UserClaim{
		UserId: activeCredential.UserId,
	})

	assert.NoError(t, err)
	assert.Equal(t, &auth_dto.UserLoginRespDto{
		SessionId: sessionId,
		Token:     token,
	}, resp)
}
