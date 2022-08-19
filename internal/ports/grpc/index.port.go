package portsGRPC

import (
	pb "togo/pkg/grpc"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHealthcheckServer
	pb.UnimplementedTaskReaderServiceServer
	pb.UnimplementedTaskWriteServiceServer
	pb.UnimplementedUserReaderServiceServer
	pb.UnimplementedUserWriteServiceServer
}

// NewServer creates a new gRPC server.
func NewServer() *grpc.Server {
	serverGRPC := grpc.NewServer()

	pb.RegisterHealthcheckServer(serverGRPC, &Server{})
	pb.RegisterTaskReaderServiceServer(serverGRPC, &Server{})
	pb.RegisterTaskWriteServiceServer(serverGRPC, &Server{})
	pb.RegisterUserReaderServiceServer(serverGRPC, &Server{})
	pb.RegisterUserWriteServiceServer(serverGRPC, &Server{})

	return serverGRPC
}
