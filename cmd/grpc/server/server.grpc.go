package cmdGRPCServer

import (
	"flag"
	"fmt"
	"log"
	"net"
	portsGRPC "togo/internal/ports/grpc"

	"github.com/go-pg/pg/v9"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func ConnectGRPCServer(db *pg.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := portsGRPC.NewServer(db)

	// reflection use for post man
	reflection.Register(server)

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
