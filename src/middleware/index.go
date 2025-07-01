package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)


func bindHeaderToContext(c *gin.Context, injectFn func(ctx context.Context, val string) context.Context, headerName string, fallbackFn func() string) {
	val := c.GetHeader(headerName)
	if val == "" {
		val = fallbackFn()
	}
	ctx := injectFn(c.Request.Context(), val)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
