# grpc2kafka
Passing on a message stream via GRPC to Kafka

Preinstalling required packages:
go get google.golang.org/grpc\n
go get -u github.com/golang/protobuf/protoc-gen-go

Creating the protbuf file:
protoc -I event/ event/event.proto --go_out=plugins=grpc:event

Installing the event as a package itself:
go get github.com/dbrandenburg/grpc2kafka/event
