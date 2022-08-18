package cmdGRPCServer

import (
	"flag"
	"fmt"
	"log"
	"net"
	portsGRPC "togo/internal/ports/grpc"
	pb "togo/pkg/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedHealthcheckServer
}

// func (s *server) GetPing(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
// 	result := servicesHealthCheck.GetPing()
// 	timestamp := timestamppb.New(result.Date)
// 	log.Printf("Received: %v", result)
// 	return &pb.PingResponse{Url: result.URL, Date: timestamp}, nil
// }

func ConnectGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := portsGRPC.NewServer()

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
