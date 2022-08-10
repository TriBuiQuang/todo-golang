package portsRestFulUser

import (
	"net/http"
	adapterPostgresRepo "togo/internal/adapter/postgressql/repositories"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func NewUserPort(db *pg.DB) *SUserPort {
	userService := &serviceUsers.UserService{
		DB:       db,
		UserRepo: &adapterPostgresRepo.SUserRepo{},
	}

	return &SUserPort{
		UserService: userService,
	}
}

type SUserPort struct {
	UserService interface {
		// GetAllUsers(c *gin.Context)
		CreateUser(user *domain.SUser) (*domain.SUser, error)
		// GetSingleUser(c *gin.Context)
		// EditUser(c *gin.Context)
		// DeleteUser(c *gin.Context)
	}
}

// func GetAllUsers(c *gin.Context) {
// 	service := &serviceUsers.UserService{}

// 	users, totalUsers, err := service.GetAllUsers()

// 	if err != nil {

// 		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":     http.StatusOK,
// 		"message":    "All Users",
// 		"total_user": totalUsers,
// 		"data":       users,
// 	})
// }

func (u *SUserPort) CreateUser(c *gin.Context) {
	user := &domain.SUser{}
	c.BindJSON(user)

	newUser, err := u.UserService.CreateUser(user)
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

// func GetSingleUser(c *gin.Context) {
// 	userId := c.Param("userId")
// 	service := &serviceUsers.UserService{
// 		User: &domain.SUser{ID: userId},
// 	}

// 	err := service.GetSingleUser()

// 	if err != nil {
// 		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusNotFound))
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Single User",
// 		"data":    service.User,
// 	})

// }

// func EditUser(c *gin.Context) {
// 	userId := c.Param("userId")
// 	service := &serviceUsers.UserService{
// 		User: &domain.SUser{ID: userId},
// 	}
// 	c.BindJSON(&service.User)
// 	// completed := user.Completed

// 	err := service.EditUser()

// 	if err != nil {
// 		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "User Edited Successfully",
// 	})

// }

// func DeleteUser(c *gin.Context) {
// 	userId := c.Param("userId")

// 	service := &serviceUsers.UserService{
// 		User: &domain.SUser{ID: userId},
// 	}
// 	err := service.DeleteUser()

// 	if err != nil {
// 		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "User deleted successfully",
// 	})
// }
