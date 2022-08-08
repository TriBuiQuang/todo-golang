package main

import (
	"log"
	cmdRestFulRoutes "togo/cmd/restful"
	adapterPostgres "togo/internal/adapter/postgressql"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect postgres database
	adapterPostgres.ConnectDatabase()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	cmdRestFulRoutes.Routes(router)

	log.Fatal(router.Run("localhost:8080"))
}
