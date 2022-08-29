package service

import (
	"http_server/pkg/repository"
)

type ApiService struct {
	repo repository.Balance
}

func newApiService(repo repository.Balance) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) CreateUser(user microservice.UsersBalances) (int, error) {

	return s.repo.CreateUser(user)
}

func (s *ApiService) GetBalanceById(userId int, ccy string) (microservice.UsersBalances, error) {

	return s.repo.GetBalanceById(userId, ccy)
}
