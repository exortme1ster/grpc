package main

import (
	"log"
	"net"

	"github.com/exortme1ster/grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
}
