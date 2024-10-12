package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppPayload struct {
	OwnerID     string `validate:"required"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Environment string `validate:"required,valid_env"`
	OwnerType   string `validate:"required,valid_owner_type"`
}

type App struct {
	ID          primitive.ObjectID `bson:"_id"`
	OwnerID     string             `bson:"ownerId"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Environment string             `bson:"environment"`
	OwnerType   string             `bson:"ownerType"`
	CreatedAt   time.Time          `bson:"createdAt"`
	LastUpdated time.Time          `bson:"lastUpdate"`
}

func NewApp(payload *AppPayload) *App {
	return &App{
		ID:          primitive.NewObjectID(),
		OwnerID:     payload.OwnerID,
		Name:        payload.Name,
		Description: payload.Description,
		Environment: payload.Environment,
		OwnerType:   payload.OwnerType,
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
}
