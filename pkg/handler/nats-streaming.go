package handler

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type ConfigNatsStreaming struct {
	Host      string
	Port      string
	ClusterID string
	ClientID  string
}

func NewNatsStreamingConnection(cfg ConfigNatsStreaming) (stan.Conn, error) {
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(fmt.Sprintf("nats://%v:%v", cfg.Host, cfg.Port)))
	if err != nil {
		return sc, fmt.Errorf("can't connect to Nats Streaming channel (subscription): %s", err)
	}
	return sc, nil
}
