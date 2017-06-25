package main

import (
  "log"
  "os"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/dbrandenburg/grpc2kafka/event"
)

const (
  address     = "localhost:50051"
  defaultName = "world"
)

func main() {
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewEventProtosClient(conn)

  name := defaultName
  if len(os.Args) > 1 {
    name = os.Args[1]
  }
  r, err := c.SendEvent(context.Background(), &pb.Event{Name: name, Latitude: 1, Longitude: 2})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Greeting: %s", r.Status)
}
