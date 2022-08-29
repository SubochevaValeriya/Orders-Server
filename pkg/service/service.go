package service

import (
	"http_server/pkg/repository"
)

type Balance interface {
	CreateUser(user microservice.UsersBalances) (int, error)
	GetBalanceById(userId int, ccy string) (microservice.UsersBalances, error)
}

type Service struct {
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newApiService(repos.Balance)}
}
