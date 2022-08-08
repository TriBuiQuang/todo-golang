package servicesHealthCheck

import (
	"time"
)

type pingResponse struct {
	URL  string    `json:"url"`
	Date time.Time `json:"date"`
}

func GetPing() pingResponse {

	return pingResponse{
		URL:  "/api/ping",
		Date: time.Now(),
	}
}
