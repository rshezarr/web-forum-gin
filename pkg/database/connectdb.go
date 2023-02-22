package database

import (
	"forum/internal/config"
	mgo "forum/pkg/database/mongodb"
	psql "forum/pkg/database/postgres"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseConnection interface {
	ConnectDB(cfg *config.Config) (*sqlx.DB, *mongo.Database, error)
}

func DetectDatabase(cfg *config.Config) (*sqlx.DB, *mongo.Database, error) {
	var factory DatabaseConnection

	switch cfg.Database.DBName {
	case "postgres":
		factory = &psql.PostgresConnectionFactory{
			Host:         viper.GetString("postgres.host"),
			Port:         viper.GetInt("postgres.port"),
			DatabaseName: viper.GetString("postgres.database_name"),
			Username:     viper.GetString("postgres.username"),
			Password:     viper.GetString("postgres.password"),
		}
	case "mongo":
		factory = &mgo.MongoConnectionFactory{
			Host:         viper.GetString("mongo.host"),
			Port:         viper.GetInt("mongo.port"),
			DatabaseName: viper.GetString("mongo.database_name"),
			Username:     viper.GetString("mongo.username"),
			Password:     viper.GetString("mongo.password"),
		}
	default:
		log.Fatalf("Invalid database type specified: %v", cfg.Database.DBName)
	}

	db, session, err := factory.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, session, nil
}
