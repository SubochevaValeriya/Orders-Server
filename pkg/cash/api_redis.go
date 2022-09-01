package cash

import (
	"github.com/go-redis/redis"
	order "http_server"
	"time"
)

type ApiCash struct {
	redisClient redis.Client
}

func newApiCash(client redis.Client) *ApiCash {
	return &ApiCash{redisClient: client}
}

func (c *ApiCash) CreateOrder(order order.Order) (int, error) {
	err = c.redisClient.Set(ctx, "products", cachedProducts, 10*time.Second).Err()

	return c.repo.CreateOrder(order)
}

func (c *ApiCash) GetOrderById(orderId int) (order.Order, error) {

	return c.repo.GetOrderById(orderId)
}
