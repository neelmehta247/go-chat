package grpc

import (
	"context"

	pb "github.com/neelmehta247/go-chat/proto"
)

type grpcServer struct{}

func NewServer() *grpcServer {
	return &grpcServer{}
}

func (s *grpcServer) SayMessage(ctx context.Context, h *pb.Hello) (*pb.Reponse, error) {
	return &pb.Reponse{
		Message: h.Message,
	}, nil
}
