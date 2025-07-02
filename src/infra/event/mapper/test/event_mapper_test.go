package mapper_test

//go:generate mockgen -destination=../../../../mock/gen/proto/kafka_event_mock.go -package=proto_mock MVC_DI/gen/proto KafkaEventServiceClient

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/context_key"
	"MVC_DI/global/enum"
	"MVC_DI/global/model"
	"MVC_DI/infra/event/mapper/impl"
	"context"
	"errors"
	"fmt"
	"testing"

	proto_mock "MVC_DI/mock/gen/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func mockContext() context.Context {
	ctx := context.Background()
	ctx = context_key.WithIdempotencyKey(ctx, "idempotency-key")
	ctx = context_key.WithCorrelationId(ctx, "correlation-id")
	ctx = context_key.WithRequestId(ctx, "request-id")
	return ctx
}
func Test_SubmitEvent_GRPCError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	mapper := &impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	innerErr := errors.New("grpc failed")
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(nil, innerErr)

	err := mapper.SubmitEvent(mockContext(), envelope)
	assert.EqualError(t, err, model.NewAppError().
		WithStatusKey(enum.SYSTEM_ERROR{}).
		WithDetail(fmt.Sprintf("non-gRPC error: %v", innerErr)).Error())
}

func Test_SubmitEvent_Async_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	mapper := &impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_ASYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{Status: proto.EventStatus_PROCESSED_SUCCESS}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.NoError(t, err)
}

func Test_SubmitEvent_Sync_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	mapper := &impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{TriggerModeEffective: proto.TriggerMode_SYNC, Status: proto.EventStatus_PROCESSED_SUCCESS}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.NoError(t, err)
}

func Test_SubmitEvent_Sync_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	mapper := &impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{Status: proto.EventStatus_FINAL_FAILED}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), enum.CODE_EVENT_FALLBACKED)
}

func Test_SubmitEvent_UnknownTriggerMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := proto_mock.NewMockKafkaEventServiceClient(ctrl)
	mapper := &impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode(999)}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{}, nil)

	defer func() {
		if r := recover(); r != nil {
			if msg, ok := r.(string); ok {
				assert.Contains(t, msg, "illegal trigger mode: 999")
			} else {
				t.Errorf("unexpected panic type: %v", r)
			}
		} else {
			t.Errorf("expected panic but did not panic")
		}
	}()

	_ = mapper.SubmitEvent(context.Background(), envelope)
}
