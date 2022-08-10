package main

import (
	"log"
	cmdRestFulRoutes "togo/cmd/restful"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect postgres database
	// db := adapterPostgres.ConnectDatabase()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	cmdRestFulRoutes.Routes(router)

	log.Fatal(router.Run("localhost:8080"))
}
