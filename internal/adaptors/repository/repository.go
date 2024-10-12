package repository

import (
	"card-validator-apps-service/internal/config"
	"card-validator-apps-service/internal/helpers"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const appsCollection = "apps"

type AppRepository interface {
}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(cfg *config.Config) (AppRepository, error) {
	ctx, cancel := helpers.NewTimeoutContext(cfg.DefaultTimeout)
	defer cancel()

	client, err := connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Client:     client,
		Collection: client.Database(cfg.DatabaseName).Collection(appsCollection),
	}, nil
}

func connect(ctx context.Context, DBURL string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBURL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
