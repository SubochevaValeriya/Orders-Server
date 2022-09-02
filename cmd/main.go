package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	server "http_server"
	"http_server/pkg/handler"
	"http_server/pkg/repository"
	"http_server/pkg/service"
	"os"
)

const (
	Port        = ":4222"
	ClusterID   = "test-cluster"
	ClientID    = "test-client"
	ChannelName = "order-channel"
)

//sudo docker run --name=orders -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
//docker run --name redis-test-instance -p 6379:6379 -d redis
//sudo docker run -p 4222:4222 -p 8223:8223 nats-streaming -p 4222 -m 8223
// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
//sudo docker exec -it e8c0fc42a2f9 /bin/bash
//psql -U postgres

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.ConfigPostgres{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	cache, err := repository.NewRedisDB(repository.ConfigRedis{
		Host:     viper.GetString("cache.host_cache"),
		Port:     viper.GetString("cache.port"),
		Password: "",
		DBName:   viper.GetInt("cache.dbname"),
	})

	if err != nil {
		logrus.Fatalf("failed to inititalize cache: %s", err.Error())
	}

	natsStreaming, err := handler.NewNatsStreamingConnection(handler.ConfigNatsStreaming{
		Host:      viper.GetString("nats-streaming.host_nats"),
		Port:      viper.GetString("nats-streaming.port"),
		ClusterID: viper.GetString("nats-streaming.cluster_id"),
		ClientID:  viper.GetString("nats-streaming.client_id"),
	})

	if err != nil {
		logrus.Fatalf("failed to connect to nats-streaming: %s", err.Error())
	}

	// dependency injection
	repos := repository.NewRepository(db, cache)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, natsStreaming)
	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
