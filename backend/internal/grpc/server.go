package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/neelmehta247/go-chat/backend/internal/grpc/service"
	pb "github.com/neelmehta247/go-chat/proto"
)

func NewGrpcServer(reflectionEnabled bool) *grpc.Server {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloWorldServer(grpcServer, service.NewHelloService())

	if reflectionEnabled {
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)
	}

	return grpcServer
}
