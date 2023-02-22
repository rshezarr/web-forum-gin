package database

import (
	"fmt"
	"forum/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostgresConnectionFactory struct {
	Host         string
	Port         int
	DatabaseName string
	Username     string
	Password     string
}

func (p *PostgresConnectionFactory) ConnectDB(cfg *config.Config) (*sqlx.DB, *mongo.Database, error) {
	logrus.Info("Connecting to database...")

	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.DatabaseURL)
	if err != nil {
		return nil, nil, fmt.Errorf("error while database connection: %w", err)
	}

	logrus.Info("Creating tables...")

	if err := goose.Up(db.DB, cfg.Database.SchemePath, goose.WithNoVersioning()); err != nil {
		logrus.Fatalf("migration: %s", err.Error())
	}

	logrus.Info("Tables are created and successful conenction")

	return db, nil, nil
}
