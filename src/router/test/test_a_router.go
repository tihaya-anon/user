package test_router

import (
	"MVC_DI/middleware"
	"MVC_DI/router"
  test_controller "MVC_DI/section/test/controller"
	"MVC_DI/util"

	"github.com/gin-gonic/gin"
)

func BindTestAController (ctrl *test_controller.TestAController) {
  router.RegisterRouter(func(publicRouterGroup *gin.RouterGroup, authRouterGroup *gin.RouterGroup) {
    publicGroup := util.RoutesWrapper(publicRouterGroup.Group("/test-a"))
    authGroup := util.RoutesWrapper(authRouterGroup.Group("/test-a"))
    authGroup.Use(middleware.JwtMiddleware())

    publicGroup.GET("/hello", ctrl.Hello)
    authGroup.GET("/hello", ctrl.Hello)
  })
}