package event

import (
	"MVC_DI/section/auth/dto"
	"context"
)

type AuthEventPublisher interface {
	PublishLoginAudit(ctx context.Context, dto *dto.PublishLoginAuditDto) error
	PublishInvalidSession(ctx context.Context, sessionId int64) error
}
