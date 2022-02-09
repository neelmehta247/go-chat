package service

import (
	"context"
	"fmt"

	pb "github.com/neelmehta247/go-chat/proto"
)

type helloService struct{}

func NewHelloService() *helloService {
	return &helloService{}
}

func (s *helloService) SayHello(ctx context.Context, h *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{
		Message: fmt.Sprintf("Hello, %s", h.GetName()),
	}, nil
}

func (s *helloService) SayHelloAgain(ctx context.Context, h *pb.HelloRequest) (*pb.HelloReponse, error) {
	return &pb.HelloReponse{
		Message: fmt.Sprintf("Hello again, %s", h.GetName()),
	}, nil
}
