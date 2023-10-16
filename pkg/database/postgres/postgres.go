package database

import (
	"fmt"
	"forum/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func ConnectDB(cfg *config.Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error while database connection: %w", err)
	}

	if err := goose.Up(db.DB, cfg.Database.SchemePath, goose.WithNoVersioning()); err != nil {
		logrus.Fatalf("migration: %s", err.Error())
	}

	logrus.Info("Tables are created and successfully connected")

	return db, nil
}
