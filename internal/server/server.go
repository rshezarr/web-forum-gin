package server

import (
	"context"
	"forum/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
	Shutdown(ctx context.Context) error
	Notify() <-chan error
}

type server struct {
	server          *http.Server
	ServerErrorChan chan error
}

func NewServer(cfg *config.Configuration, router *gin.Engine) Server {
	return &server{
		server: &http.Server{
			Addr:           cfg.API.Addr,
			Handler:        router,
			MaxHeaderBytes: cfg.API.MaxHeaderBytes << 20,
			ReadTimeout:    cfg.API.ReadTimeout,
			WriteTimeout:   cfg.API.WriteTimeout,
			IdleTimeout:    cfg.API.IdleTimeout,
		},
		ServerErrorChan: make(chan error, 1),
	}
}

func (s *server) Run() {
	s.ServerErrorChan <- s.server.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *server) Notify() <-chan error {
	return s.ServerErrorChan
}
