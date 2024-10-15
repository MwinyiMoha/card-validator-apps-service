package application

import (
	"card-validator-apps-service/internal/adaptors/repository"
	"card-validator-apps-service/internal/core/domain"

	"github.com/go-playground/validator/v10"
)

type AppService interface {
	CreateApp(payload *domain.AppPayload) (*domain.App, error)
	ListApps(userId string) ([]*domain.App, error)
	FetchApp(userId, appId string) (*domain.App, error)
	DeleteApp(userId, appId string) error
	RefreshKey(userId, appId string) (string, error)
	DecodeKey(appKey string) (string, error)
}

type Service struct {
	Repository repository.AppRepository
	Validator  *validator.Validate
}

func NewService(repo repository.AppRepository, val *validator.Validate) AppService {
	return &Service{
		Repository: repo,
		Validator:  val,
	}
}
