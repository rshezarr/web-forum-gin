package mongodb

import (
	"context"
	"fmt"
	"forum/internal/config"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnectionFactory struct {
	Host         string
	Port         int
	DatabaseName string
	Username     string
	Password     string
}

func (factory *MongoConnectionFactory) ConnectDB(cfg *config.Config) (*sqlx.DB, *mongo.Database, error) {
	cliOptions := options.Client().ApplyURI("mongodb://mongo_db:27017")

	client, err := mongo.Connect(context.Background(), cliOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("connect: %w", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, nil, fmt.Errorf("ping: %w", err)
	}

	return nil, client.Database("user-service"), nil
}
