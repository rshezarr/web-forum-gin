package app

import (
	"context"
	"forum/internal/config"
	"forum/internal/handler"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"forum/pkg/database"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, _, err := database.DetectDatabase(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	server := server.NewServer(cfg, handler.InitRoutes())

	go func() {
		logrus.Fatal("error while running server: %s", server.Run().Error())
	}()

	logrus.Info("Forum app starting...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Forum app shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Error("error while server shutting down: %s", err.Error())
	}

	logrus.Info("Closing database...")

	if err := db.Close(); err != nil {
		logrus.Error("error while closing database: %s", err.Error())
	}
}
