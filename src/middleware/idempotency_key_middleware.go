package middleware

import (
	"MVC_DI/global/context_key"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const IdempotencyKeyHeader string = "X-Idempotency-Key"

func IdempotencyKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bindHeaderToContext(c, context_key.WithIdempotencyKey, IdempotencyKeyHeader, func() string { return uuid.NewString() })
	}
}
