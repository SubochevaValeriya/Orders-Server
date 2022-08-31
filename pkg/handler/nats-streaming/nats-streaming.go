package handler

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	Port        = ":4222"
	ClusterID   = "test-cluster"
	ClientID    = "test-client"
	ChannelName = "order-channel"
)

//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223

//Subscription to a channel and receiving data from it

func Subscription() ([]byte, error) {
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(fmt.Sprintf("nats://%v%v", viper.GetString("nats-streaming.host_nats"), Port)))
	if err != nil {
		logrus.Fatalf("can't connect to Nats Streaming channel (subscription): %s", err)
	}

	var msg []byte

	sub, err := sc.Subscribe(ChannelName,
		func(m *stan.Msg) {
			msg = m.Data
			m.Ack()
		},
		stan.StartWithLastReceived())

	err = sub.Close()

	if err != nil {
		return nil, fmt.Errorf("error in Nats streaming subscription: %w", err)
	}
	fmt.Println(string(msg), "from nats")
	return msg, nil
}
