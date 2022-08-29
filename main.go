package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

const (
	//port      = ":50051"
	port      = ":8223"
	clusterID = "test-cluster"
	clientID  = "event-store"
)

//sudo docker run -p 4223:4223 -p 8223:8223 nats-streaming -p 4223 -m 8223

func main() {
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

//func main() {
//	//	s, err := stand.RunServer("mystreamingserver")
//
//	opts := stand.GetDefaultOptions()
//	opts.StoreType = stores.TypeFile
//	opts.FilestoreDir = "datastore"
//	s, err := stand.RunServerWithOpts(opts, nil)
//
//	fmt.Println(s, err)
//}

//type server struct{}
//
//// CreateOrder RPC creates a new Event into EventStore
//func (s *server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
//	// Persist data into EventStore database
//	command := store.EventStore{}
//	err := command.CreateEvent(in)
//	if err != nil {
//		return nil, err
//	}
//	// Publish event on NATS Streaming Server
//	go publishEvent(in)
//	return &pb.Response{IsSuccess: true}, nil
//}
//
//func (s *server) GetEvents(ctx context.Context, in *pb.EventFilter) (*pb.EventResponse, error) {
//	return &pb.EventResponse{Events: nil}, nil
//}
//
//// publishEvent publish an event via NATS Streaming server
//func publishEvent(event *pb.Event) {
//	// Connect to NATS Streaming server
//	sc, err := stan.Connect(
//		clusterID,
//		clientID,
//		stan.NatsURL(stan.DefaultNatsURL),
//	)
//	if err != nil {
//		log.Print(err)
//		return
//	}
//	defer sc.Close()
//	channel := event.Channel
//	eventMsg := []byte(event.EventData)
//	// Publish message on subject (channel)
//	sc.Publish(channel, eventMsg)
//	log.Println("Published message on channel: " + channel)
//}
//
//func main() {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	// Creates a new gRPC server
//	s := grpc.NewServer()
//	pb.RegisterEventStoreServer(s, &server{})
//	s.Serve(lis)
//}

//func hash(){
//	m := map[int]struct{}
//}
