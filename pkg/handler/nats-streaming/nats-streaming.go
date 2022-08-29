package handler

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

const (
	port      = ":4223"
	clusterID = "test-cluster"
	clientID  = "event-store"
)

//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223

func subscription() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost:4223"))
	fmt.Println(err)
	err = sc.Publish("foo", []byte("Hello World")) //запаблишую с другого клиента
	fmt.Println(err)
	sub, err := sc.Subscribe("foo",
		func(m *stan.Msg) {
			fmt.Println(string(m.Data))
			m.Ack()
		},
		stan.StartWithLastReceived())
	err = sub.Close()
	fmt.Println(err, sub)
}
