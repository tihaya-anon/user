package user_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
  user_controller "MVC_DI/section/user/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindUserController (ctrl *user_controller.UserController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/user"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/user"))
    authGroup.Use(middleware.JwtMiddleware())
  })
}