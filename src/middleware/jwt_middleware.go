package middleware

import (
	"MVC_DI/global/enum"
	"MVC_DI/security"
	"MVC_DI/vo/resp"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := resp.NewResponse()
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.MISSING_TOKEN, enum.MSG.MISSING_TOKEN, nil))
			return
		}
		isLegal, token := extractToken(token)
		if !isLegal || !security.CheckJWT(token) {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.INVALID_TOKEN, enum.MSG.INVALID_TOKEN, nil))
			return
		}
		ctx.Set("token", &token)
		ctx.Next()
	}
}
func GetToken(ctx *gin.Context) *string {
	token, exists := ctx.Get("token")
	if !exists {
		return nil
	}
	str, ok := token.(*string)
	if !ok {
		return nil
	}
	return str
}
func extractToken(token string) (bool, string) {
	isLegal := strings.HasPrefix(token, "Bearer ") && len(strings.Split(token, " ")) == 2
	if isLegal {
		return true, strings.Split(token, " ")[1]
	}
	return false, ""
}
