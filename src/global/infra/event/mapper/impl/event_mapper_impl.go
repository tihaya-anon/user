package event_mapper_impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/enum"
	global_model "MVC_DI/global/model"
	event_mapper "MVC_DI/global/infra/event/mapper"
	"context"
)

type EventMapperImpl struct {
	KafkaEventServiceClient proto.KafkaEventServiceClient
}

// SubmitEvent implements event_mapper.EventMapper.
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
			return global_model.NewAppError().WithCode(enum.CODE.GRPC_ERROR).WithDetail(response.GetStatus().String())
		}
		return nil

	default:
		return global_model.NewAppError().WithCode(enum.CODE.UNKNOWN_TRIGGER_MODE).WithMessage(enum.MSG.UNKNOWN_TRIGGER_MODE).WithDetail(envelope.GetTriggerModeRequested().String())
	}
}

// INTERFACE
var _ event_mapper.EventMapper = (*EventMapperImpl)(nil)
