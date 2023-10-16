package server

import (
	"context"
	"forum/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server          *http.Server
	ServerErrorChan chan error
}

func NewServer(cfg *config.Configuration, router *gin.Engine) *Server {
	return &Server{
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

func (s *Server) Run() {
	s.ServerErrorChan <- s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.ServerErrorChan
}
