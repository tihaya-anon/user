package auth_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
	"MVC_DI/section/auth/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindAuthController(ctrl *controller.AuthController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		// publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/auth"))
		authGroup := util.RoutesWrapper(authRouterGroup.Group("/auth"))
		authGroup.Use(middleware.JwtMiddleware())
	})
}
