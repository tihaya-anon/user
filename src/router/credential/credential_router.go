package credential_router

import (
	"MVC_DI/router"
	credential_controller "MVC_DI/section/credential/controller"

	"github.com/gin-gonic/gin"
)

func BindCredentialController(ctrl *credential_controller.CredentialController) {
	router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
		// publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/credential"))
		// authGroup := router.RoutesWrapper(authRouterGroup.Group("/credential"))
	})
}
