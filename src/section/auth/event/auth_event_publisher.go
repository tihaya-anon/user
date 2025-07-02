package event

import (
	"MVC_DI/section/auth/dto"
	"context"
)

//go:generate mockgen -source=auth_event_publisher.go -destination=..\..\..\mock\auth\event\auth_event_publisher_mock.go -package=event_mock
type AuthEventPublisher interface {
	PublishLoginAudit(ctx context.Context, dto *dto.PublishLoginAuditDto) error
	PublishInvalidSession(ctx context.Context, sessionId int64) error
}
