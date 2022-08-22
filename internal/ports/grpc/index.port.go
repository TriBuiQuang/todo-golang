package portsGRPC

import (
	pb "togo/pkg/grpc"

	"github.com/go-pg/pg/v9"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHealthcheckServer
}

// NewServer creates a new gRPC server.
func NewServer(db *pg.DB) *grpc.Server {
	serverGRPC := grpc.NewServer()

	userPort := NewUserPort(db)
	taskPort := NewTaskPort(db)

	pb.RegisterHealthcheckServer(serverGRPC, &Server{})
	pb.RegisterTaskReaderServiceServer(serverGRPC, taskPort)
	pb.RegisterTaskWriteServiceServer(serverGRPC, taskPort)
	pb.RegisterUserReaderServiceServer(serverGRPC, userPort)
	pb.RegisterUserWriteServiceServer(serverGRPC, userPort)

	return serverGRPC
}
