package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	server "github.com/neelmehta247/go-chat/backend/internal/grpc"
)

func main() {
	port := flag.Int("port", 8443, "port to run grpc server on")
	reflectionOn := flag.Bool("reflection", true, "turn on server reflection")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := server.NewGrpcServer(reflectionOn != nil && *reflectionOn)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Running server on port %d\n", *port)
}
