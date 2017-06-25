package main

import (
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/dbrandenburg/grpc2kafka/event"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) SendEvent(ctx context.Context, in *pb.Event) (*pb.Delivery, error) {
  return &pb.Delivery{Status: "Hello " + in.Name}, nil
}

func (s *server) GetEvents(loc *pb.Location, stream pb.EventProtos_GetEventsServer) error {
  //return &pb.Event{Latitude: 1, Longitude: 2}, nil
  return nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterEventProtosServer(s, &server{})
  s.Serve(lis)
}
