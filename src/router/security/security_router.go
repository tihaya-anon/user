package security_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
  security_controller "MVC_DI/section/security/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindSecurityController (ctrl *security_controller.SecurityController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/security"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/security"))
    authGroup.Use(middleware.JwtMiddleware())
  })
}