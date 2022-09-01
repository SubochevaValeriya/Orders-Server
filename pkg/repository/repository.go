package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	order "http_server"
)

type Order interface {
	CreateOrder(order order.Order) (int, error)
	GetOrderById(orderId int) (order.Order, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB, cash *redis.Client) *Repository {
	return &Repository{NewApiPostgres(db, cash)}
}
