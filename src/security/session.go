package security

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSessionId(ctx *gin.Context) *int64 {
	sessionIdStr, err := ctx.Cookie("sessionId")
	if err != nil {
		return nil
	}
	sessionId, err := strconv.ParseInt(sessionIdStr, 10, 64)
	if err != nil {
		return nil
	}
	return &sessionId
}

func SetSessionId(ctx *gin.Context, sessionId int64, maxAge int, path string, domain string, secure bool, httpOnly bool) {
	ctx.SetCookie("sessionId", strconv.FormatInt(sessionId, 10), maxAge, path, domain, secure, httpOnly)
}
