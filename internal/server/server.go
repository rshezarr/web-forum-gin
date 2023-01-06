package server

import (
	"context"
	"forum/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		server: &http.Server{
			Addr:           cfg.API.Addr,
			Handler:        router,
			MaxHeaderBytes: cfg.API.MaxHeaderBytes << 20,
			ReadTimeout:    cfg.API.ReadTimeout,
			WriteTimeout:   cfg.API.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	logrus.Infof("Server is started at port http://localhost%s", s.server.Addr)

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Server shutting down...")

	return s.server.Shutdown(ctx)
}
