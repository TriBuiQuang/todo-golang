package cmdGRPCClient

import (
	"context"
	"flag"
	"log"
	"time"
	portsGRPC "togo/internal/ports/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func ConnectGRPCClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := portsGRPC.NewHealthcheckClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetPing(ctx, &portsGRPC.PingRequest{})
	if err != nil {
		log.Fatalf("could not ping: %v", err)
	}
	log.Printf("Ping: %s - %s", r.GetUrl(), r.GetDate())

}
