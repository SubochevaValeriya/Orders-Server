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

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//sudo docker run --name=balance -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
	// migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

	//sudo docker exec -it e8c0fc42a2f9 /bin/bash
	//psql -U postgres

	//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223
	db, err := repository.NewPostgresDB(repository.Config{
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

	// dependency injection
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	//sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(fmt.Sprintf("nats://%v%v", viper.GetString("db.host_nats"), Port)))
	//if err != nil {
	//	logrus.Fatalf("can't connect to Nats Streaming channel (subscription): %s", err)
	//}
	//
	//startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	//
	//subj, i := ":4222", 0
	//mcb := func(msg *stan.Msg) {
	//	i++
	//	println(msg, i)
	//}
	//
	//_, err = sc.QueueSubscribe(subj, "qgroup", mcb, startOpt, stan.DurableName(""))
	//if err != nil {
	//	sc.Close()
	//	log.Fatal(err)
	//}
	//
	//log.Printf("Listening o")
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//var tpl *template.Template
//
//func templates() error {
//	tpl, err := template.ParseGlob("/home/valeriya/Документы/GitHub/http-server/templates/index.templates")
//	if err != nil {
//		return fmt.Errorf("error while parse templates template: %w", err)
//	}
//
//	tpl.ExecuteTemplate()
//
//}
//
//func searchHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("***searchHandler is running***")
//	tpl.ExecuteTemplate(w, "index.templates", nil)
//}
