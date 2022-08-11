package portsRestFul

import (
	"net/http"
	"togo/internal/core/domain"

	"github.com/gin-gonic/gin"
)

/*
Validate receive json data , if doesn't have name and user_id return true, otherwise return false
*/
func ValidateCreateTask(c *gin.Context, task *domain.STask) bool {
	if task.Name == "" || task.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Name or user_id must not empty",
		})
		return true
	}
	return false
}
