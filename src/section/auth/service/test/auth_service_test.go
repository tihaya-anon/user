package auth_service_test

import (
	"MVC_DI/gen/proto"
	global_model "MVC_DI/global/model"
	auth_mapper_mock "MVC_DI/mock/auth/mapper"
	auth_service_mock "MVC_DI/mock/auth/service"
	event_mapper_mock "MVC_DI/mock/event/mapper"
	auth_dto "MVC_DI/section/auth/dto"
	auth_enum "MVC_DI/section/auth/enum"
	auth_service_impl "MVC_DI/section/auth/service/impl"
	"MVC_DI/security/jwt/claims"
	"MVC_DI/security/jwt"
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
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
	assert.EqualError(t, err, global_model.NewAppError().WithCode(auth_enum.CODE.UNKNOWN_CREDENTIAL).WithMessage(auth_enum.MSG.UNKNOWN_CREDENTIAL).Error())
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

	token, err := jwt.GenerateJWT(claims.UserClaim{
		UserId: activeCredential.UserId,
	})

	assert.NoError(t, err)
	assert.Equal(t, &auth_dto.UserLoginRespDto{
		SessionId: sessionId,
		Token:     token,
	}, resp)
}
func Test_LogoutUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventMapper := event_mapper_mock.NewMockEventMapper(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		EventMapper: mockEventMapper,
	}

	ctx := &gin.Context{}
	sessionId := int64(101)
	envelope := &proto.KafkaEnvelope{
		DeliveryMode:         proto.DeliveryMode_PUSH,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
		Payload:              []byte(strconv.FormatInt(sessionId, 10)),
	}

	mockEventMapper.EXPECT().
		SubmitEvent(ctx, envelope).
		Return(nil)

	err := svc.LogoutUser(ctx, sessionId)
	assert.NoError(t, err)
}

func Test_LogoutUser_SubmitEventError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventMapper := event_mapper_mock.NewMockEventMapper(ctrl)
	svc := auth_service_impl.AuthServiceImpl{
		EventMapper: mockEventMapper,
	}

	ctx := &gin.Context{}
	sessionId := int64(202)
	envelope := &proto.KafkaEnvelope{
		DeliveryMode:         proto.DeliveryMode_PUSH,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
		Payload:              []byte(strconv.FormatInt(sessionId, 10)),
	}

	simulatedErr := errors.New("event mapper failure")
	mockEventMapper.EXPECT().
		SubmitEvent(ctx, envelope).
		Return(simulatedErr)

	err := svc.LogoutUser(ctx, sessionId)
	assert.EqualError(t, err, simulatedErr.Error())
}
