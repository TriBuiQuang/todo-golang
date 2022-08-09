package portsRestFulUser

import (
	"net/http"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
)

type SUserPort struct {
	UserInterface interface {
		GetAllUsers(c *gin.Context)
		CreateUser(c *gin.Context)
		GetSingleUser(c *gin.Context)
		EditUser(c *gin.Context)
		DeleteUser(c *gin.Context)
	}
}

func GetAllUsers(c *gin.Context) {
	service := &serviceUsers.UserService{}

	users, totalUsers, err := service.GetAllUsers()

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
	service := &serviceUsers.UserService{
		User: &domain.SUser{},
	}
	c.BindJSON(&service.User)

	newUser, err := service.CreateUser()

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
	service := &serviceUsers.UserService{
		User: &domain.SUser{ID: userId},
	}

	err := service.GetSingleUser()

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    service.User,
	})

}

func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	service := &serviceUsers.UserService{
		User: &domain.SUser{ID: userId},
	}
	c.BindJSON(&service.User)
	// completed := user.Completed

	err := service.EditUser()

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

	service := &serviceUsers.UserService{
		User: &domain.SUser{ID: userId},
	}
	err := service.DeleteUser()

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
}
