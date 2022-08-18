package portsGRPC

import (
	pb "togo/pkg/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHealthcheckServer
}

// NewServer creates a new gRPC server.
func NewServer() *grpc.Server {

	serverGRPC := grpc.NewServer()

	pb.RegisterHealthcheckServer(serverGRPC, &Server{})

	return serverGRPC
}
