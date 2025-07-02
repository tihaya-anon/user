package router

import (
	"MVC_DI/router"
  "MVC_DI/section/permission/controller"

	"github.com/gin-gonic/gin"
)

func BindPermissionController (ctrl *controller.PermissionController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := router.RoutesWrapper(publicRouterGroup)
    // authGroup := router.RoutesWrapper(authRouterGroup)
  })
}