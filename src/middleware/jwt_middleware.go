package middleware

import (
	"MVC_DI/global/context_key"
	"MVC_DI/global/enum"
	"MVC_DI/vo/resp"
	"strings"

	"github.com/gin-gonic/gin"
)

const JwtHeader string = "Authorization"
const JwtPrefix string = "Bearer "

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := resp.NewResponse()
		token := ctx.GetHeader(JwtHeader)
		if token == "" {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE_MISSING_TOKEN, enum.MSG_MISSING_TOKEN, nil))
			return
		}
		isLegal, token := extractToken(token)
		if !isLegal {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE_INVALID_TOKEN, enum.MSG_INVALID_TOKEN, nil))
			return
		}
		c := context_key.WithJwt(ctx.Request.Context(), token)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
func extractToken(token string) (bool, string) {
	isLegal := strings.HasPrefix(token, JwtPrefix) && len(strings.Split(token, " ")) == 2
	if isLegal {
		return true, strings.Split(token, " ")[1]
	}
	return false, ""
}
