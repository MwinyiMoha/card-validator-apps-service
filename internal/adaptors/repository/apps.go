package repository

import (
	"card-validator-apps-service/internal/core/domain"
	"card-validator-apps-service/internal/helpers"
	"context"

	"github.com/mwinyimoha/card-validator-utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const readTimeout = 5

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.Client.Disconnect((ctx))
}

func (r *Repository) SaveApp(app *domain.App) error {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, app)
	if err != nil {
		return errors.WrapError(err, errors.Internal, "failed to save document")
	}

	return nil
}

func (r *Repository) FindApps(filter bson.M) ([]*domain.App, error) {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	findOptions := options.Find().SetSort(bson.M{"createdAt": -1})
	cursor, err := r.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, errors.WrapError(err, errors.Internal, "internal error")
	}

	apps := make([]*domain.App, 0)
	if err := cursor.All(ctx, &apps); err != nil {
		return nil, errors.WrapError(err, errors.Internal, "internal error")
	}

	return apps, nil
}

func (r *Repository) FindApp(filter bson.M) (*domain.App, error) {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	return r.FindAppOrError(ctx, filter)
}

func (r *Repository) DeleteApp(filter bson.M) error {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	_, err := r.FindAppOrError(ctx, filter)
	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.WrapError(err, errors.Internal, "failed to delete record")
	}

	return nil
}

func (r *Repository) RefreshKey(filter bson.M, newKey string) (string, error) {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	app, err := r.FindAppOrError(ctx, filter)
	if err != nil {
		return "", err
	}

	app.AppKey = newKey
	_, err = r.Collection.ReplaceOne(ctx, filter, app)
	if err != nil {
		return "", errors.WrapError(err, errors.Internal, "failed to update record")
	}

	return newKey, nil
}

func (r *Repository) ValidateKey(key string) (string, error) {
	ctx, cancel := helpers.NewTimeoutContext(readTimeout)
	defer cancel()

	app, err := r.FindAppOrError(ctx, bson.M{"appKey": key})
	if err != nil {
		return "", err
	}

	return app.ID.Hex(), nil
}

func (r *Repository) FindAppOrError(ctx context.Context, filter bson.M) (*domain.App, error) {
	result := r.Collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.WrapErrorf(err, errors.NotFound, "record not found")
		}

		return nil, errors.WrapErrorf(err, errors.Internal, "internal error")
	}

	app := new(domain.App)
	if err := result.Decode(app); err != nil {
		return nil, errors.WrapErrorf(err, errors.Internal, "internal error")
	}

	return app, nil
}
