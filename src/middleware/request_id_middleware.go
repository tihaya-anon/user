package middleware

import (
	"MVC_DI/global/context_key"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIdHeader string = "X-Request-ID"

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bindHeaderToContext(c, context_key.WithRequestId, RequestIdHeader, func() string { return uuid.NewString() })
	}
}
