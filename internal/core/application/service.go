package application

import "card-validator-apps-service/internal/adaptors/repository"

type AppService interface {
}

type Service struct {
	Repository repository.AppRepository
}

func NewService(repo repository.AppRepository) AppService {
	return &Service{
		Repository: repo,
	}
}
