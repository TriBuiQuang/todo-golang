package portsRestFulUser

import (
	"errors"
	"net/http"
	adapterPostgresRepo "togo/internal/adapter/postgresql/repositories"
	"togo/internal/core/domain"
	serviceUsers "togo/internal/core/services/users"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

type SUserPort struct {
	UserService interface {
		CreateUser(user *domain.SUser) (*domain.SUser, error)
		GetAllUsers() ([]domain.SUser, int, error)
		GetSingleUser(user *domain.SUser) error
		EditUser(user *domain.SUser) error
		DeleteUser(user *domain.SUser) error
	}
}

func NewUserPort(db *pg.DB) *SUserPort {
	userService := &serviceUsers.UserService{
		DB: db,
		UserRepo: &adapterPostgresRepo.SUserRepo{
			DB: db,
		},
	}

	return &SUserPort{
		UserService: userService,
	}
}

func (u *SUserPort) CreateUser(c *gin.Context) {
	user := &domain.SUser{}
	err := c.BindJSON(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusForbidden))
		return
	}

	if user.Username == "" {
		c.JSON(portsRestFul.PrintErrResponse(errors.New("this req is missing username "), http.StatusBadRequest))
		return
	}

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

func (u *SUserPort) GetAllUsers(c *gin.Context) {
	users, totalUsers, err := u.UserService.GetAllUsers()

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

func (u *SUserPort) GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &domain.SUser{ID: userId}

	err := u.UserService.GetSingleUser(user)

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

func (u *SUserPort) EditUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &domain.SUser{ID: userId}

	c.BindJSON(user)

	if user.Username == "" && user.LimitPerDay == 0 {
		c.JSON(portsRestFul.PrintErrResponse(errors.New("must have username or limit_per_day value"), http.StatusBadRequest))
		return
	}

	err := u.UserService.EditUser(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User Edited Successfully",
	})
}

func (u *SUserPort) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &domain.SUser{ID: userId}

	err := u.UserService.DeleteUser(user)

	if err != nil {
		c.JSON(portsRestFul.PrintErrResponse(err, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
}
