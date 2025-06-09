package server

import (
	"MVC_DI/config"
	"MVC_DI/middleware"
	"MVC_DI/router"
	"context"
	"errors"
	"fmt"
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
	authRouterGroup.Use(middleware.JwtMiddleware())

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
			fmt.Printf("Start server failed: %s\n", err)
			return
		}
	}()
}

func (s *Server) Stop(timeOut time.Duration) {
	ctx, cancelShutdown := context.WithTimeout(context.Background(), timeOut)
	defer cancelShutdown()

	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
		return
	}
	fmt.Println("Server shutdown successfully")
}
