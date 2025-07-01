package router

import (
	"MVC_DI/middleware"
	"MVC_DI/vo/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type THandlerFunc = func(ctx *gin.Context) *resp.TResponse

type TRoutesWrapper struct {
	routes      gin.IRoutes
	method      string
	path        string
	middlewares []gin.HandlerFunc
}
type tMiddlewareWrapper struct {
	r *TRoutesWrapper
}

func RoutesWrapper(routes gin.IRoutes) *TRoutesWrapper {
	return &TRoutesWrapper{
		routes: routes,
	}
}

func iHandler2GinHandler(fn []THandlerFunc) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, len(fn))
	for i, fn := range fn {
		handlers[i] = resp.HandlerWrapper(fn)
	}
	return handlers
}
func (w *tMiddlewareWrapper) Idem() *tMiddlewareWrapper {
	w.r.middlewares = append(w.r.middlewares, middleware.IdempotencyKeyMiddleware())
	return w
}
func (w *tMiddlewareWrapper) Handler(handler ...THandlerFunc) {
	all := append(w.r.middlewares, iHandler2GinHandler(handler)...)
	w.r.routes.Handle(w.r.method, w.r.path, all...)
}

func (w *TRoutesWrapper) GET(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodGet
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) POST(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodPost
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) DELETE(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodDelete
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) PATCH(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodPatch
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) PUT(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodPut
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) OPTIONS(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodOptions
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
func (w *TRoutesWrapper) HEAD(relativePath string) *tMiddlewareWrapper {
	w.method = http.MethodHead
	w.path = relativePath
	return &tMiddlewareWrapper{w}
}
