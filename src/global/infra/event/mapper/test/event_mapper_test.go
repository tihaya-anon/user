package event_mapper_test

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/enum"
	event_mapper_impl "MVC_DI/global/infra/event/mapper/impl"
	"context"
	"errors"
	"testing"

	mock_proto "MVC_DI/mock/gen/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_SubmitEvent_GRPCError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_proto.NewMockKafkaEventServiceClient(ctrl)
	mapper := &event_mapper_impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(nil, errors.New("grpc failed"))

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.EqualError(t, err, "grpc failed")
}

func Test_SubmitEvent_Async_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_proto.NewMockKafkaEventServiceClient(ctrl)
	mapper := &event_mapper_impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_ASYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{Status: proto.EventStatus_PROCESSED_SUCCESS}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.NoError(t, err)
}

func Test_SubmitEvent_Sync_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_proto.NewMockKafkaEventServiceClient(ctrl)
	mapper := &event_mapper_impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{Status: proto.EventStatus_PROCESSED_SUCCESS}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.NoError(t, err)
}

func Test_SubmitEvent_Sync_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_proto.NewMockKafkaEventServiceClient(ctrl)
	mapper := &event_mapper_impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode_SYNC}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{Status: proto.EventStatus_FINAL_FAILED}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), enum.CODE.GRPC_ERROR)
}

func Test_SubmitEvent_UnknownTriggerMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_proto.NewMockKafkaEventServiceClient(ctrl)
	mapper := &event_mapper_impl.EventMapperImpl{KafkaEventServiceClient: mockClient}

	envelope := &proto.KafkaEnvelope{TriggerModeRequested: proto.TriggerMode(999)}
	mockClient.EXPECT().SubmitEvent(gomock.Any(), gomock.Any()).Return(&proto.SubmitEventResponse{}, nil)

	err := mapper.SubmitEvent(context.Background(), envelope)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), enum.CODE.UNKNOWN_TRIGGER_MODE)
}
