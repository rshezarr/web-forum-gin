package app

import (
	"context"
	"forum/internal/config"
	"forum/internal/http/v1"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	database "forum/pkg/database/postgres"
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

	db, err := database.ConnectDB(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := v1.NewController(svc)
	srv := server.NewServer(cfg, ctrl.InitRoutes())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		logrus.Infof("Server is started at port: %s", cfg.API.Addr)
		srv.Run()
	}()

	select {
	case sig := <-quit:
		logrus.Info("app: signal accepted: %s", sig.String())
	case err := <-srv.ServerErrorChan:
		logrus.Info("app: signal accepted: %s", err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Error("error while server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error while closing database: %s", err.Error())
	}
}
