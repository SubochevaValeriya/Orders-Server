package service

import (
	order "http_server"
	"http_server/pkg/repository"
)

type Order interface {
	CreateOrder(order order.Order) (int, error)
	GetOrderById(orderId int) (order.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{newApiService(repos.Balance)}
}
