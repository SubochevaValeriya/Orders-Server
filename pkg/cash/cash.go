package cash

import (
	"github.com/jmoiron/sqlx"
	order "http_server"
	"http_server/pkg/repository"
)

type Order interface {
	CreateOrder(order order.Order) (int, error)
	GetOrderById(orderId int) (order.Order, error)
}

type Cash struct {
	Order
}

func NewCash(repos *repository.Repository) *Cash {
	return &Cash{newCashInMemory(),
}
