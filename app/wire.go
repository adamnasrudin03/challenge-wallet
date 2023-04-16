package app

import (
	"adamnasrudin03/challenge-wallet/app/repository"
	"adamnasrudin03/challenge-wallet/app/service"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		User:        repository.NewUserRepository(db),
		Transaction: repository.NewTransactionRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		User: service.NewUserService(repo.User),
	}
}
