package main

import (
  "log"
  "net"

  pb "event/event"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) Process(ctx context.Context, in *pb.Event) (*pb.Delivery, error) {
  return &pb.Event{Message: "Hello "}, nil
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
