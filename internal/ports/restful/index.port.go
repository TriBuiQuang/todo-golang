package portsRestFul

import (
	"fmt"
	"log"
	servicesHealthCheck "togo/internal/core/services/healthcheck"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}

func HealthCheck(c *gin.Context) {
	fmt.Println("data gin.context: ", c.Request.URL)
	ping := servicesHealthCheck.GetPing()

	c.IndentedJSON(http.StatusOK, ping)
}

/**  Only handle print error message:
 *	- 1xx: Informational - Request received, continuing process
 *	- 2xx: Success - The action was successfully received, understood, and accepted
 *	- 3xx: Redirection - Further action must be taken in order to complete the request
 *	- 4xx: Client Error - The request contains bad syntax or cannot be fulfilled
 *	- 5xx: Server Error - The server failed to fulfill an apparently valid request
 */
func PrintErrResponse(err error, httpStatus int) (int, map[string]any) {
	var errMessage string
	errStatusCode := httpStatus

	log.Printf("Error - Reason: %v\n", err)
	switch errStatusCode {
	case 100:
		errMessage = "Information - Something went wrong"
	case 200:
		errMessage = "Success - Something went wrong"
	case 300:
		errMessage = "Redirection - Something went wrong"
	case 400:
		errMessage = "Client error - Something went wrong"
	case 500:
		errMessage = "Server Error - Something went wrong"
	default:
		errMessage = "Don't have this status code- everything good"
	}

	errResponse := map[string]any{
		"status":  errStatusCode,
		"message": errMessage,
	}

	return errStatusCode, errResponse
}
