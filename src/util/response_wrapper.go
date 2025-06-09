package util

import (
	"MVC_DI/vo/resp"

	"github.com/gin-gonic/gin"
)

type IHandlerFunc = func(ctx *gin.Context) *resp.TResponse

type IRoutesWrapper struct {
	Routes gin.IRoutes
}

func RoutesWrapper(routes gin.IRoutes) *IRoutesWrapper {
	return &IRoutesWrapper{
		Routes: routes,
	}
}

func (wrapper *IRoutesWrapper) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return wrapper.Routes.Use(middleware...)
}

func iHandler2GinHandler(fn []IHandlerFunc) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, len(fn))
	for i, fn := range fn {
		handlers[i] = resp.HandlerWrapper(fn)
	}
	return handlers
}

func (wrapper *IRoutesWrapper) GET(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.GET(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) POST(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.POST(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) DELETE(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.DELETE(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) PATCH(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.PATCH(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) PUT(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.PUT(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) OPTIONS(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.OPTIONS(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (wrapper *IRoutesWrapper) HEAD(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return wrapper.Routes.HEAD(relativePath, iHandler2GinHandler(responseFunc)...)
}
