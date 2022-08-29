package repository

import (
	"github.com/jmoiron/sqlx"
	order "http_server"
)

type Balance interface {
	CreateOrder(order order.Order) (int, error)
	GetOrderById(orderId int) (order.Order, error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewApiPostgres(db)}
}
