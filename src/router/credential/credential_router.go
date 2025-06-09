package credential_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
  credential_controller "MVC_DI/section/credential/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindCredentialController (ctrl *credential_controller.CredentialController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/credential"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/credential"))
    authGroup.Use(middleware.JwtMiddleware())
  })
}