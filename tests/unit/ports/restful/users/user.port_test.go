package portsRestFulUser

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"togo/internal/core/domain"
	portsRestFulUser "togo/internal/ports/restful/users"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func MockJsonPost(c *gin.Context /* the test context */, content interface{}) {
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
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

func TestCreateUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name         string
		req          *domain.SUser
		sendMessage  map[string]interface{}
		setup        func()
		expectedCode int
		expectedErr  error
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
			expectedCode: http.StatusCreated,
			expectedErr:  nil,
		},
		{
			name: "fail because not give anything",
			req:  &domain.SUser{},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(nil),
				}
			},
			expectedCode: http.StatusBadRequest,
			expectedErr:  nil,
		},
		{
			name: "fail because not give Username",
			req:  &domain.SUser{LimitPerDay: 2},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(nil),
				}
			},
			expectedCode: http.StatusBadRequest,
			expectedErr:  nil,
		},
		{
			name: "fail because give empty json",
			req:  &domain.SUser{},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: handleCreateUser(errors.New("test fail")),
				}
			},
			expectedCode: http.StatusBadRequest,
			expectedErr:  nil,
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
			expectedCode: http.StatusInternalServerError,
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(response)
			c.Request = &http.Request{
				Header: make(http.Header),
			}

			MockJsonPost(c, testCase.req)
			userPort.CreateUser(c)

			assert.Equal(t, testCase.expectedCode, response.Code)
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name string

		sendMessage  map[string]interface{}
		setup        func()
		expectedCode int
		expectedErr  error
	}{
		{
			name: "success",
			setup: func() {
				userService = &mockUserService{}
			},
			expectedCode: http.StatusOK,
			expectedErr:  nil,
		},
		{
			name: "fail because fail in service layer",
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			expectedCode: http.StatusInternalServerError,
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(response)
			c.Request = &http.Request{
				Header: make(http.Header),
			}

			userPort.GetAllUsers(c)

			assert.Equal(t, testCase.expectedCode, response.Code)
		})
	}
}

func TestSingleUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name         string
		params       []gin.Param
		sendMessage  map[string]interface{}
		setup        func()
		expectedCode int
		expectedErr  error
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
			expectedCode: http.StatusOK,
			expectedErr:  nil,
		},
		{
			name:   "fail because fail in service layer",
			params: []gin.Param{},
			setup: func() {
				userService = &mockUserService{
					err: errors.New("test fail"),
				}
			},
			expectedCode: http.StatusNotFound,
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setup()
			userPort := &portsRestFulUser.SUserPort{
				UserService: userService,
			}

			// Mock gin package
			response := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(response)
			c.Request = &http.Request{
				Header: make(http.Header),
			}
			c.Params = testCase.params
			userPort.GetSingleUser(c)

			assert.Equal(t, testCase.expectedCode, response.Code)
		})
	}
}
