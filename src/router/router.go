package router

import (
	"github.com/gin-gonic/gin"
)

type IRegisterRouterFunc = func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup)

var RegisterRouterFuncList []IRegisterRouterFunc

func RegisterRouter(fn IRegisterRouterFunc) {
	if fn == nil {
		return
	}
	RegisterRouterFuncList = append(RegisterRouterFuncList, fn)
}
