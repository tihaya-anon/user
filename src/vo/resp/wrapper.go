package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerWrapper(fn func(ctx *gin.Context) *TResponse) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ResponseWrapper(ctx, fn(ctx))
	}
}

func ResponseWrapper(ctx *gin.Context, response *TResponse) {
	if response == nil {
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}
