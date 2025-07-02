package impl

import (
	"MVC_DI/gen/proto"
	"MVC_DI/infra/event/envelope"
	"MVC_DI/infra/event/mapper"
	"MVC_DI/section/auth/dto"
	"MVC_DI/section/auth/event"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthEventPublisherImpl struct {
	EventMapper          mapper.EventMapper
	EventEnvelopeFactory envelope.EventEnvelopeFactory
}

// PublishInvalidSession implements event.AuthEventPublisher.
func (a *AuthEventPublisherImpl) PublishInvalidSession(ctx context.Context, sessionId int64) error {
	// TODO: dynamically decide the trigger mode
	message := &proto.InvalidateSessionRequest{SessionId: sessionId}
	eventEnvelope, err := a.EventEnvelopeFactory.Build(
		&envelope.EventSubmissionDto{
			Message:      message,
			DeliveryMode: proto.DeliveryMode_PUSH,
			Priority:     proto.Priority_HIGH,
			TriggerMode:  proto.TriggerMode_ASYNC,
		},
	)
	if err != nil {
		return err
	}
	eventEnvelope.TopicName = envelope.AUTH_EVENT_TOPIC
	eventEnvelope.EventType = event.INNVALID_SESSION_EVENT
	err = a.EventMapper.SubmitEvent(ctx, eventEnvelope)
	return err
}

// PublishLoginAudit implements event.AuthEventPublisher.
func (a *AuthEventPublisherImpl) PublishLoginAudit(ctx context.Context, dto *dto.PublishLoginAuditDto) error {
	message := &proto.AddLoginAuditLogRequest{
		UserId:     dto.UserId,
		LoginTime:  timestamppb.Now(),
		IpAddress:  dto.IpAddress,
		DeviceInfo: dto.DeviceInfo,
		Result:     dto.Result,
	}
	eventEnvelope, err := a.EventEnvelopeFactory.Build(
		&envelope.EventSubmissionDto{
			Message:      message,
			DeliveryMode: proto.DeliveryMode_PULL,
			Priority:     proto.Priority_LOW,
			TriggerMode:  proto.TriggerMode_ASYNC,
		},
	)
	if err != nil {
		return err
	}
	eventEnvelope.TopicName = envelope.AUTH_EVENT_TOPIC
	eventEnvelope.EventType = event.LOGIN_AUDIT_EVENT
	err = a.EventMapper.SubmitEvent(ctx, eventEnvelope)
	return err
}

// INTERFACE
var _ event.AuthEventPublisher = (*AuthEventPublisherImpl)(nil)
