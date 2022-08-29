package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

const (
	Port              = ":4223"
	ClusterID         = "test-cluster"
	ClientID          = "test-client"
	ChannelName       = "order-channel"
	DataForPublishing = "\"{\\\"order_uid\\\":\\\"b563feb7b2b84b6test\\\",\\\"track_number\\\":\\\"WBILMTESTTRACK\\\",\\\"entry\\\":\\\"WBIL\\\",\\\"delivery\\\":{\\\"name\\\":\\\"TestTestov\\\",\\\"phone\\\":\\\"+9720000000\\\",\\\"zip\\\":\\\"2639809\\\",\\\"city\\\":\\\"KiryatMozkin\\\",\\\"address\\\":\\\"PloshadMira15\\\",\\\"region\\\":\\\"Kraiot\\\",\\\"email\\\":\\\"test@gmail.com\\\"},\\\"payment\\\":{\\\"transaction\\\":\\\"b563feb7b2b84b6test\\\",\\\"request_id\\\":\\\"\\\",\\\"currency\\\":\\\"USD\\\",\\\"provider\\\":\\\"wbpay\\\",\\\"amount\\\":1817,\\\"payment_dt\\\":1637907727,\\\"bank\\\":\\\"alpha\\\",\\\"delivery_cost\\\":1500,\\\"goods_total\\\":317,\\\"custom_fee\\\":0},\\\"items\\\":[{\\\"chrt_id\\\":9934930,\\\"track_number\\\":\\\"WBILMTESTTRACK\\\",\\\"price\\\":453,\\\"rid\\\":\\\"ab4219087a764ae0btest\\\",\\\"name\\\":\\\"Mascaras\\\",\\\"sale\\\":30,\\\"size\\\":\\\"0\\\",\\\"total_price\\\":317,\\\"nm_id\\\":2389212,\\\"brand\\\":\\\"VivienneSabo\\\",\\\"status\\\":202}],\\\"locale\\\":\\\"en\\\",\\\"internal_signature\\\":\\\"\\\",\\\"customer_id\\\":\\\"test\\\",\\\"delivery_service\\\":\\\"meest\\\",\\\"shardkey\\\":\\\"9\\\",\\\"sm_id\\\":99,\\\"date_created\\\":\\\"2021-11-26T06:22:19Z\\\",\\\"oof_shard\\\":\\\"1\\\"}\""
)

//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223

func main() {
	publishing()
}

//publishing data to a channel

func publishing() {
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(fmt.Sprintf("nats://localhost%v", Port)))
	if err != nil {
		logrus.Fatalf("can't connect to Nats Streaming channel (publishing): %s", err)
	}

	err = sc.Publish(ChannelName, []byte(DataForPublishing))

	if err != nil {
		logrus.Fatalf("can't publish data to channel: %s", err)
	}

	logrus.Println("data is published")
}
