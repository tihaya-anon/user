package middleware

import (
	"MVC_DI/global/context_key"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CorrelationIdHeader string = "X-Correlation-ID"

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.GetHeader(CorrelationIdHeader)
		if cid == "" {
			cid = uuid.New().String()
		}
		ctx := context_key.WithCorrelationId(c.Request.Context(), cid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
