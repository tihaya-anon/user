package server

import (
	"MVC_DI/config"
	"MVC_DI/global"
	"MVC_DI/middleware"
	"MVC_DI/router"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Setup(publicPath, authPath string, engine *gin.Engine) {
	publicRouterGroup := engine.Group(publicPath)
	authRouterGroup := engine.Group(authPath)

	publicRouterGroup.Use(middleware.CorrelationIdMiddleware())
	publicRouterGroup.Use(middleware.RequestIdMiddleware())

	authRouterGroup.Use(middleware.JwtMiddleware())
	authRouterGroup.Use(middleware.CorrelationIdMiddleware())
	authRouterGroup.Use(middleware.RequestIdMiddleware())

	for _, registerRouterFunc := range router.RegisterRouterFuncList {
		registerRouterFunc(publicRouterGroup, authRouterGroup)
	}

	s.server = &http.Server{
		Addr:    config.Application.App.Uri,
		Handler: engine,
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Red.Printf("Server Start failed: %s\n", err)
			return
		}
	}()
}

func (s *Server) Stop(timeOut time.Duration) {
	ctx, cancelShutdown := context.WithTimeout(context.Background(), timeOut)
	defer cancelShutdown()

	if err := s.server.Shutdown(ctx); err != nil {
		global.Red.Printf("Server shutdown failed: %s\n", err)
		return
	}
	global.Green.Println("Server shutdown successfully")
}
