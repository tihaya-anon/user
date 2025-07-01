package router

import (
	"MVC_DI/router"
  "MVC_DI/section/credential/controller"

	"github.com/gin-gonic/gin"
)

func BindCredentialController (ctrl *controller.CredentialController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/credential"))
    // authGroup := router.RoutesWrapper(authRouterGroup.Group("/credential"))
  })
}