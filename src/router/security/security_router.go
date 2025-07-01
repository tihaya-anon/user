package router

import (
	"MVC_DI/router"
  "MVC_DI/section/security/controller"

	"github.com/gin-gonic/gin"
)

func BindSecurityController (ctrl *controller.SecurityController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/security"))
    // authGroup := router.RoutesWrapper(authRouterGroup.Group("/security"))
  })
}