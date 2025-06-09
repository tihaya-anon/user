package permission_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
  permission_controller "MVC_DI/section/permission/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindPermissionController (ctrl *permission_controller.PermissionController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/permission"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/permission"))
    authGroup.Use(middleware.JwtMiddleware())
  })
}