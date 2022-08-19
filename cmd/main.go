package main

import (
	"flag"
	cmdGRPCClient "togo/cmd/grpc/client"
	cmdGRPCServer "togo/cmd/grpc/server"
	cmdRestFulRoutes "togo/cmd/restful"
	adapterPostgres "togo/internal/adapter/postgresql"
)

func main() {
	// Get arg from command line
	wordArgCommand := flag.String("protocol", "default value", "a string for description")
	flag.Parse()
	db := adapterPostgres.ConnectDatabase()

	switch *wordArgCommand {
	case "server":
		cmdGRPCServer.ConnectGRPCServer(db)
		return
	case "client":
		cmdGRPCClient.ConnectGRPCClient()
		return
	default:
		cmdRestFulRoutes.ConnectRestFull()
		return

	}

}
