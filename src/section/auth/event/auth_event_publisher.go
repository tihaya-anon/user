package auth_event_publisher

import (
	auth_dto "MVC_DI/section/auth/dto"
	"context"
)

type AuthEventPublisher interface {
	PublishLoginAudit(ctx context.Context, dto *auth_dto.PublishLoginAuditDto) error
	PublishInvalidSession(ctx context.Context, sessionId int64) error
}
