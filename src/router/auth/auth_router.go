package auth_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
	auth_controller "MVC_DI/section/auth/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindAuthController(ctrl *auth_controller.AuthController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		// publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/auth"))
		authGroup := util.RoutesWrapper(authRouterGroup.Group("/auth"))
		authGroup.Use(middleware.JwtMiddleware())
	})
}
