package router

import (
	"MVC_DI/router"
  "MVC_DI/section/user/controller"

	"github.com/gin-gonic/gin"
)

func BindUserController (ctrl *controller.UserController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    // publicGroup := router.RoutesWrapper(publicRouterGroup.Group("/user"))
    // authGroup := router.RoutesWrapper(authRouterGroup.Group("/user"))
  })
}