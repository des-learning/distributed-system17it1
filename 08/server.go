package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	"des.com/hellogrpc/hello"
)

type helloServer struct{}

func (s *helloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	n := rand.Int63n(6)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return &hello.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	sock, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("error listening to network: %v\n", err)
	}

	server := grpc.NewServer()
	hello.RegisterHelloServer(server, &helloServer{})
	rand.Seed(time.Now().Unix())
	err = server.Serve(sock)
	if err != nil {
		log.Fatalf("error starting grpc server: %v\n", err)
	}
}
