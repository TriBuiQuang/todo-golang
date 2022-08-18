package portsGRPC

import (
	"context"
	"log"
	servicesHealthCheck "togo/internal/core/services/healthcheck"
	pb "togo/pkg/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetPing(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	result := servicesHealthCheck.GetPing()
	timestamp := timestamppb.New(result.Date)
	log.Printf("Received: %v", result)
	return &pb.PingResponse{Url: result.URL, Date: timestamp}, nil
}
