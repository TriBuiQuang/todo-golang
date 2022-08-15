package main

import (
	"log"

	cmdRestFulRoutes "togo/cmd/restful"
	adapterPostgres "togo/internal/adapter/postgressql"
)

func main() {
	// Connect postgres database
	// db := adapterPostgres.ConnectDatabase()

	// Init Router
	db := adapterPostgres.ConnectDatabase()

	router := cmdRestFulRoutes.NewRoot(db)
	log.Fatal(router.Run("localhost:8080"))
}
