package cmd

import (
	"MVC_DI/config"
	"MVC_DI/global"
	"MVC_DI/server"
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func bindController() {

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
	global.Grey.Println("============= START =============")
	bindController()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	global.Green.Printf("activate profile : %s\n", config.Application.Env)
	global.Green.Printf("listen to        : %s\n", config.Application.App.Uri)
	publicPath := "/api/v1/public"
	authPath := "/api/v1/auth"
	global.Green.Printf("public path      : %s\n", publicPath)
	global.Green.Printf("auth path        : %s\n", authPath)
	startServer(publicPath, authPath, engine, 5*time.Second)
}

func Stop() {
	global.Grey.Println("============= STOP ==============")
}
