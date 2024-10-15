package application

import (
	"card-validator-apps-service/internal/core/domain"
	"card-validator-apps-service/internal/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/mwinyimoha/card-validator-utils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const keyLength = 8

func (s *Service) CreateApp(payload *domain.AppPayload) (*domain.App, error) {
	if err := s.Validator.Struct(payload); err != nil {
		if verr, ok := err.(validator.ValidationErrors); ok {
			violations := errors.BuildViolations(verr)
			return nil, errors.NewValidationError(violations)
		}

		return nil, err
	}

	code, err := helpers.NewRandomCode(keyLength)
	if err != nil {
		return nil, err
	}

	app := domain.NewApp(payload, code)
	if err := s.Repository.SaveApp(app); err != nil {
		return nil, err
	}

	return app, nil
}

func (s *Service) ListApps(userId string) ([]*domain.App, error) {
	filter := map[string]interface{}{
		"ownerId": userId,
	}
	return s.Repository.FindApps(filter)
}

func (s *Service) FetchApp(userId, appId string) (*domain.App, error) {
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return nil, errors.WrapError(err, errors.BadRequest, "invalid app id")
	}

	filter := map[string]interface{}{
		"_id":     id,
		"ownerId": userId,
	}
	return s.Repository.FindApp(filter)
}

func (s *Service) DeleteApp(userId, appId string) error {
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return errors.WrapError(err, errors.BadRequest, "invalid app id")
	}

	filter := map[string]interface{}{
		"_id":     id,
		"ownerId": userId,
	}
	return s.Repository.DeleteApp(filter)
}

func (s *Service) RefreshKey(userId, appId string) (string, error) {
	id, err := primitive.ObjectIDFromHex(appId)
	if err != nil {
		return "", errors.WrapError(err, errors.BadRequest, "invalid app id")
	}

	code, err := helpers.NewRandomCode(keyLength)
	if err != nil {
		return "", err
	}

	filter := map[string]interface{}{
		"_id":     id,
		"ownerId": userId,
	}
	return s.Repository.RefreshKey(filter, code)
}

func (s *Service) DecodeKey(appKey string) (string, error) {
	return s.Repository.ValidateKey(appKey)
}
