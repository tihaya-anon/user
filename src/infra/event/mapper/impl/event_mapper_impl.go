package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/enum"
	"MVC_DI/global/model"
	"MVC_DI/infra/event/mapper"
	"context"
)

type EventMapperImpl struct {
	KafkaEventServiceClient proto.KafkaEventServiceClient
}

// SubmitEvent implements mapper.EventMapper.
func (e *EventMapperImpl) SubmitEvent(ctx context.Context, envelope *proto.KafkaEnvelope) error {
	response, err := e.KafkaEventServiceClient.SubmitEvent(ctx, &proto.SubmitEventRequest{Envelope: envelope})
	if err != nil {
		return err
	}

	switch envelope.GetTriggerModeRequested() {
	case proto.TriggerMode_ASYNC:
		return nil

	case proto.TriggerMode_SYNC:
		if response.GetStatus() != proto.EventStatus_PROCESSED_SUCCESS {
			return model.NewAppError().WithStatusKey(enum.GRPC_ERROR{}).WithDetail(response.GetStatus().String())
		}
		return nil

	default:
		return model.NewAppError().WithStatusKey(enum.UNKNOWN_TRIGGER_MODE{}).WithDetail(envelope.GetTriggerModeRequested().String())
	}
}

// INTERFACE
var _ mapper.EventMapper = (*EventMapperImpl)(nil)
