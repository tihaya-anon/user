package cmd

import (
	"MVC_DI/config"
	"MVC_DI/global/log"
	test_router "MVC_DI/router/test"
	test_controller_builder "MVC_DI/section/test/controller/builder"
	test_mapper_builder "MVC_DI/section/test/mapper/builder"
	test_service_builder "MVC_DI/section/test/service/builder"
	"MVC_DI/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func bindController() {
	testAMapper := test_mapper_builder.NewTestAMapperBuilder().Build()
	testAService := test_service_builder.NewTestAServiceBuilder().WithTestAMapper(testAMapper).Build()
	testAController := test_controller_builder.NewTestAControllerBuilder().WithTestAService(testAService).Build()
	testAController.Logger = log.GetLogger(24 * time.Hour)
	test_router.BindTestAController(testAController)
}

func startServer(publicPath, authPath string, engine *gin.Engine, timeOut time.Duration) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	newServer := server.NewServer()
	newServer.Setup(publicPath, authPath, engine)
	newServer.Run()

	<-ctx.Done()

	newServer.Stop(timeOut)
}

func Start() {
	fmt.Println("============= START =============")
	bindController()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	fmt.Printf("activate profile: %s\n", config.Application.Env)
	fmt.Printf("listen to: %s\n", config.Application.App.Uri)
	publicPath := "/api/v1/public"
	authPath := "/api/v1/auth"
	fmt.Printf("public path: %s\n", publicPath)
	fmt.Printf("auth path: %s\n", authPath)
	startServer(publicPath, authPath, engine, 5*time.Second)
}

func Stop() {
	fmt.Println("============= STOP =============")
}
