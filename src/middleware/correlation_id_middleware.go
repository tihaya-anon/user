package middleware

import (
	"MVC_DI/global/context_key"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CorrelationIdHeader string = "X-Correlation-ID"

func CorrelationIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bindHeaderToContext(c, context_key.WithCorrelationId, CorrelationIdHeader, func() string { return uuid.NewString() })
	}
}
