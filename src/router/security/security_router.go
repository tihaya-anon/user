package security_router

import (
	"MVC_DI/router"
	security_controller "MVC_DI/section/security/controller"

	"github.com/gin-gonic/gin"
)

func BindSecurityController(ctrl *security_controller.SecurityController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		// publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/security"))
		// authGroup := router.RoutesWrapper(authRouterGroup.Group("/security"))
	})
}
