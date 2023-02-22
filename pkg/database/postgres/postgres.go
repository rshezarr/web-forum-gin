package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func ConnectDB(driver, databaseURL, schemePath string) (*sqlx.DB, error) {
	logrus.Info("Connecting to database...")

	db, err := sqlx.Connect(driver, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("error while database connection: %w", err)
	}

	logrus.Info("Creating tables...")

	if err := goose.Up(db.DB, schemePath, goose.WithNoVersioning()); err != nil {
		logrus.Fatalf("migration: %s", err.Error())
	}

	logrus.Info("Tables are created and successful conenction")

	return db, nil
}
