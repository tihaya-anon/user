package auth_router

import (
	"MVC_DI/router"
	"MVC_DI/section/auth/controller"

	"github.com/gin-gonic/gin"
)

func BindAuthController(ctrl *controller.AuthController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/auth"))
		authGroup := router.RoutesWrapper(authRouterGroup.Group("/auth"))

		publicGroup.POST("/login").Idem().Handler(ctrl.LoginUser)
		authGroup.POST("/logout").Idem().Handler(ctrl.LogoutUser)
	})
}
