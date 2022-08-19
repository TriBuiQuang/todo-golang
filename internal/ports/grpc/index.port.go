package portsGRPC

import (
	pb "togo/pkg/grpc"

	"github.com/go-pg/pg/v9"
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
func NewServer(db *pg.DB) *grpc.Server {
	serverGRPC := grpc.NewServer()

	// userPort := portsRestFulUser.NewUserPort(db)
	taskPort := NewTaskPort(db)

	pb.RegisterHealthcheckServer(serverGRPC, &Server{})
	pb.RegisterTaskReaderServiceServer(serverGRPC, taskPort)
	// pb.RegisterTaskWriteServiceServer(serverGRPC, &Server{})
	// pb.RegisterUserReaderServiceServer(serverGRPC, &Server{})
	// pb.RegisterUserWriteServiceServer(serverGRPC, &Server{})

	return serverGRPC
}
