package portsRestFulUser

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	mockRestFul "togo/__mocks__/ports/restful"
	"togo/internal/core/domain"
	portsRestFulUser "togo/internal/ports/restful/users"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type mockUserService struct {
	mockCreateUser func(user *domain.SUser) (*domain.SUser, error)
	err            error
}

func handleCreateUser(errMessage error) func(user *domain.SUser) (*domain.SUser, error) {
	return func(user *domain.SUser) (*domain.SUser, error) {
		if errMessage != nil {
			return &domain.SUser{ID: "random id"}, nil
		}

		return &domain.SUser{}, errMessage
	}
}

func (m *mockUserService) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	return m.mockCreateUser(user)
}

func (m *mockUserService) GetAllUsers() ([]domain.SUser, int, error) {

	return nil, -1, m.err
}

func (m *mockUserService) GetSingleUser(user *domain.SUser) error {
	return m.err
}

func (m *mockUserService) EditUser(user *domain.SUser) error {
	return m.err
}

func (m *mockUserService) DeleteUser(user *domain.SUser) error {
	return m.err
}

func TestNewTaskPort(t *testing.T) {

	result := portsRestFulUser.NewUserPort(&pg.DB{})

	assert.Equal(t, "*portsRestFulUser.SUserPort", reflect.TypeOf(result).String())
	assert.Equal(t, "*serviceUsers.UserService", reflect.TypeOf(result.UserService).String())
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name        string
		req         *domain.SUser
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",
			req: &domain.SUser{
				Username:    "test user",
				LimitPerDay: 5,
			},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(nil),
				}
			},
			actualCode: http.StatusCreated,
			actualErr:  nil,
		},
		{
			name: "fail because not give anything",
			req:  &domain.SUser{},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(nil),
				}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because not give Username",
			req:  &domain.SUser{LimitPerDay: 2},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(nil),
				}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because give empty json",
			req:  &domain.SUser{},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(errors.New("test fail")),
				}
			},
			actualCode: http.StatusBadRequest,
			actualErr:  nil,
		},
		{
			name: "fail because give normal json but fail in service layer",
			req:  &domain.SUser{Username: "testing username"},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: func(user *domain.SUser) (*domain.SUser, error) {
						return &domain.SUser{ID: "random id"}, errors.New("test fail ")
					},
				}
			},
			actualCode: http.StatusInternalServerError,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			mockRestFul.MockJsonPost(c, testCase.req)
			userPort.CreateUser(c)

			assert.Equal(t, response.Code, testCase.actualCode)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name string

		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",
			setup: func() {
				userService = &mockUserService{}
			},
			actualCode: http.StatusOK,
			actualErr:  nil,
		},
		{
			name: "fail because fail in service layer",
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusInternalServerError,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			userPort.GetAllUsers(c)

			assert.Equal(t, response.Code, testCase.actualCode)
		})
	}
}

func TestSingleUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name        string
		params      []gin.Param
		sendMessage map[string]interface{}
		setup       func()
		actualCode  int
		actualErr   error
	}{
		{
			name: "success",
			params: []gin.Param{
				{
					Key:   "userId",
					Value: "first document",
				},
			},
			setup: func() {
				userService = &mockUserService{}
			},
			actualCode: http.StatusOK,
			actualErr:  nil,
		},
		{
			name:   "fail because fail in service layer",
			params: []gin.Param{},
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			actualCode: http.StatusNotFound,
			actualErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()

			c.Params = testCase.params
			userPort.GetSingleUser(c)

			assert.Equal(t, response.Code, testCase.actualCode)
		})
	}
}

func TestEditUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name        string
		params      []gin.Param
		req         *domain.SUser
		sendMessage map[string]interface{}
		setup       func()
		expectCode  int
		expectErr   error
	}{
		{
			name: "success",
			params: []gin.Param{
				{
					Key:   "userId",
					Value: "first document",
				},
			},
			req: &domain.SUser{Username: "testing", LimitPerDay: 20},
			setup: func() {
				userService = &mockUserService{}
			},
			expectCode: http.StatusOK,
			expectErr:  nil,
		},
		{
			name: "fail because there is no username and limit per day",
			params: []gin.Param{
				{
					Key:   "userId",
					Value: "first document",
				},
			},
			req: &domain.SUser{},
			setup: func() {
				userService = &mockUserService{}
			},
			expectCode: http.StatusBadRequest,
			expectErr:  nil,
		},
		{
			name: "fail because fail in service layer",
			params: []gin.Param{{
				Key:   "userId",
				Value: "first document",
			}},
			req: &domain.SUser{Username: "testing", LimitPerDay: 20},
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			expectCode: http.StatusInternalServerError,
			expectErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()
			mockRestFul.MockJsonPost(c, testCase.req)
			c.Params = testCase.params
			userPort.EditUser(c)

			assert.Equal(t, testCase.expectCode, response.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name        string
		params      []gin.Param
		req         *domain.SUser
		sendMessage map[string]interface{}
		setup       func()
		expectCode  int
		expectErr   error
	}{
		{
			name: "success",
			params: []gin.Param{
				{
					Key:   "userId",
					Value: "first document",
				},
			},
			req: &domain.SUser{Username: "testing", LimitPerDay: 20},
			setup: func() {
				userService = &mockUserService{}
			},
			expectCode: http.StatusOK,
			expectErr:  nil,
		},
		{
			name: "fail because fail in service layer",
			params: []gin.Param{{
				Key:   "userId",
				Value: "first document",
			}},
			req: &domain.SUser{Username: "testing", LimitPerDay: 20},
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			expectCode: http.StatusInternalServerError,
			expectErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response, c := mockRestFul.MockSetup()
			// mockRestFul.MockJsonPost(c, testCase.req)
			c.Params = testCase.params
			userPort.DeleteUser(c)

			assert.Equal(t, testCase.expectCode, response.Code)
		})
	}
}
