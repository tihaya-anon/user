package user_router

import (
	"MVC_DI/router"
	user_controller "MVC_DI/section/user/controller"

	"github.com/gin-gonic/gin"
)

func BindUserController(ctrl *user_controller.UserController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		// publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/user"))
		// authGroup := router.RoutesWrapper(authRouterGroup.Group("/user"))
	})
}
