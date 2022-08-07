package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	grpc_api "com/github/stepasite/webapp-grpc/grpc/proto"
)

var (
	port = flag.Int("grpc port", 50051, "The grpc server port")
)

type server struct {
	grpc_api.UnimplementedLoginServer
}

func (s *server) Login(ctx context.Context, in *grpc_api.LoginRequest) (*grpc_api.LoginReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &grpc_api.LoginReply{Token: "Token: " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpc_api.RegisterLoginServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}