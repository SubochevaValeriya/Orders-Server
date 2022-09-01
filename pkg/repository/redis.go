package repository

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type ConfigRedis struct {
	Host     string
	Port     string
	Password string
	DBName   int
}

//Connection to Redis
//docker run --name redis-test-instance -p 6379:6379 -d redis
func NewRedisDB(cfg ConfigRedis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Password: fmt.Sprintf("%v", cfg.Password),
		DB:       cfg.DBName})

	pong, err := client.Ping().Result()

	if err != nil {
		return nil, fmt.Errorf("connect error (Redis): %s", err)
	}

	logrus.Printf("redis: succesful response: %v", pong)

	return client, nil
}
