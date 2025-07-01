package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/global/context_key"
	"MVC_DI/infra/avro/serializer"
	"MVC_DI/infra/event/mapper"
	"MVC_DI/section/auth/dto"
	"MVC_DI/section/auth/event"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthEventPublisherImpl struct {
	AvroSerializer serializer.IAvroSerializer
	EventMapper    mapper.EventMapper
}

// PublishInvalidSession implements event.AuthEventPublisher.
func (a *AuthEventPublisherImpl) PublishInvalidSession(ctx context.Context, sessionId int64) error {
	// TODO: dynamically decide the trigger mode
	addLoginAuditLogRequest := &proto.InvalidateSessionRequest{SessionId: sessionId}
	subject, schemaId, payload, err := a.AvroSerializer.SerializeProtoMessage(addLoginAuditLogRequest)
	if err != nil {
		return err
	}
	envelope := &proto.KafkaEnvelope{
		CorrelationId:        context_key.GetCorrelationId(ctx),
		SchemaSubject:        subject,
		SchemaId:             schemaId,
		Priority:             proto.Priority_HIGH,
		Payload:              payload,
		DeliveryMode:         proto.DeliveryMode_PUSH,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
	}
	err = a.EventMapper.SubmitEvent(ctx, envelope)
	if err != nil {
		// TODO log err
	}
	return err
}

// PublishLoginAudit implements event.AuthEventPublisher.
func (a *AuthEventPublisherImpl) PublishLoginAudit(ctx context.Context, dto *dto.PublishLoginAuditDto) error {
	addLoginAuditLogRequest := &proto.AddLoginAuditLogRequest{
		UserId:     dto.UserId,
		LoginTime:  timestamppb.Now(),
		IpAddress:  dto.IpAddress,
		DeviceInfo: dto.DeviceInfo,
		Result:     dto.Result,
	}
	subject, schemaId, payload, err := a.AvroSerializer.SerializeProtoMessage(addLoginAuditLogRequest)
	if err != nil {
		// TODO log err
		return nil
	}
	envelope := &proto.KafkaEnvelope{
		CorrelationId:        context_key.GetCorrelationId(ctx),
		SchemaSubject:        subject,
		SchemaId:             schemaId,
		Priority:             proto.Priority_LOW,
		Payload:              payload,
		DeliveryMode:         proto.DeliveryMode_PULL,
		TriggerModeRequested: proto.TriggerMode_ASYNC,
	}
	err = a.EventMapper.SubmitEvent(ctx, envelope)
	if err != nil {
		// TODO log error
	}
	return nil
}

// INTERFACE
var _ event.AuthEventPublisher = (*AuthEventPublisherImpl)(nil)
