package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	server "github.com/neelmehta247/go-chat/backend/_internal/grpc"
	pb "github.com/neelmehta247/go-chat/proto"
)

func main() {
	port := flag.Int("port", 8443, "port to run grpc server on")
	reflectionOn := flag.Bool("reflection", true, "turn on server reflection")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloWorldServer(grpcServer, server.NewServer())

	if reflectionOn != nil && *reflectionOn {
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Running server on port %d\n", *port)
}
