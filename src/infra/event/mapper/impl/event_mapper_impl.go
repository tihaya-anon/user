package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/context_key"
	"MVC_DI/infra/event/handler"
	"MVC_DI/infra/event/mapper"
	"MVC_DI/middleware"
	"context"
)

type EventMapperImpl struct {
	KafkaEventServiceClient proto.KafkaEventServiceClient
}

// SubmitEvent implements mapper.EventMapper.
func (e *EventMapperImpl) SubmitEvent(ctx context.Context, envelope *proto.KafkaEnvelope) error {
	envelope.IdempotencyKey = context_key.GetIdempotencyKey(ctx)
	envelope.CorrelationId = context_key.GetCorrelationId(ctx)

	envelope.Headers[middleware.IdempotencyKeyHeader] = envelope.IdempotencyKey
	envelope.Headers[middleware.CorrelationIdHeader] = envelope.CorrelationId
	envelope.Headers[middleware.RequestIdHeader] = context_key.GetRequestId(ctx)
	response, err := e.KafkaEventServiceClient.SubmitEvent(ctx, &proto.SubmitEventRequest{Envelope: envelope})
	if err != nil {
		return handler.HandleGrpcError(err)
	}
	return handler.ValidateEventResponse(envelope, response)
}

// INTERFACE
var _ mapper.EventMapper = (*EventMapperImpl)(nil)
