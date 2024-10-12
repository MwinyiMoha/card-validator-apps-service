package application

import (
	"card-validator-apps-service/internal/adaptors/repository"

	"github.com/go-playground/validator/v10"
)

type AppService interface {
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
