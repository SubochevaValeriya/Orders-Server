package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

const (
	port      = ":4223"
	clusterID = "test-cluster"
	clientID  = "event-store"
)

//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223

func main() {
	publishing()
}

func publishing() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(fmt.Sprintf("nats://localhost%v", port)))
	if err != nil {
		logrus.Fatalf("Can't connect to Nats Streaming channel (publishing): %v", err)
	}

	err = sc.Publish("foo", []byte("Hello World"))

	if err != nil {
		logrus.Fatalf("Can't publish data to channel: %v", err)
	}
}
