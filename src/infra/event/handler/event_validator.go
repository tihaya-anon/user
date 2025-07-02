package handler

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/enum"
	"MVC_DI/global/model"
	"fmt"
)

func ValidateEventResponse(
	envelope *proto.KafkaEnvelope,
	resp *proto.SubmitEventResponse,
) error {

	reqMode := envelope.GetTriggerModeRequested()
	effMode := resp.GetTriggerModeEffective()
	status := resp.GetStatus()

	switch reqMode {
	case proto.TriggerMode_ASYNC:
		return nil

	case proto.TriggerMode_SYNC:
		if effMode != proto.TriggerMode_SYNC {
			return model.NewAppError().
				WithStatusKey(enum.EVENT_FALLBACKED{}).
				WithDetail(map[string]any{
					"requested": reqMode.String(),
					"effective": effMode.String(),
					"eventId":   resp.GetEventId(),
				})
		}

		if status != proto.EventStatus_PROCESSED_SUCCESS {
			return model.NewAppError().
				WithStatusKey(enum.EVENT_FAILED{}).
				WithDetail(map[string]any{
					"status":  status.String(),
					"eventId": resp.GetEventId(),
				})
		}

		return nil

	default:
		panic(fmt.Sprintf("illegal trigger mode: %s", reqMode.String()))
	}
}
