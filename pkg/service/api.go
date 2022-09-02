package service

import (
	order "http_server"
	"http_server/pkg/repository"
)

type ApiService struct {
	repo repository.Order
}

func newApiService(repo repository.Order) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) CreateOrder(order order.Order) (int, error) {

	return s.repo.CreateOrder(order)
}

func (s *ApiService) GetOrderById(orderId int) (order.Order, error) {

	return s.repo.GetOrderById(orderId)

}
