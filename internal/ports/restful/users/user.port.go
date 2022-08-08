package portsRestFulUser

import (
	"net/http"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, totalUsers, err := serviceUsers.GetAllUsers()

	if err != nil {

		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "All Users",
		"total_user": totalUsers,
		"data":       users,
	})
}

func CreateUser(c *gin.Context) {
	var user domain.SUser
	c.BindJSON(&user)

	newUser, err := serviceUsers.CreateUser(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
		"data":    newUser,
	})
}

func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &domain.SUser{ID: userId}

	err := serviceUsers.GetSingleUser(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})

}

func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	var user domain.SUser
	c.BindJSON(&user)
	// completed := user.Completed

	err := serviceUsers.EditUser(userId, user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Edited Successfully",
	})

}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &domain.SUser{ID: userId}

	err := serviceUsers.DeleteUser(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
}
