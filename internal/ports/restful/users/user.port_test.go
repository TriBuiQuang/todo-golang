package portsRestFulUser

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"togo/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type mockUserService struct {
	mockCreateUser func(user *domain.SUser) (*domain.SUser, error)
}

// func newMockUserService() UserServiceInterface {
// 	return &mockUserService{}
// }

func (m *mockUserService) GetAllUsers() ([]domain.SUser, int, error) {
	return nil, -1, nil
}

func (m *mockUserService) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	return m.mockCreateUser(user)
}

func (s *mockUserService) GetSingleUser(user *domain.SUser) error {
	return nil
}

func (s *mockUserService) EditUser(user *domain.SUser) error {
	return nil
}

func (s *mockUserService) DeleteUser(user *domain.SUser) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	userService := &mockUserService{}
	testCases := []struct {
		name         string
		req          *domain.SUser
		setup        func()
		expectedCode int
		expectedErr  error
	}{
		{
			name: "create user success",
			req: &domain.SUser{
				Username:    "test user",
				LimitPerDay: 5,
			},
			setup: func() {
				userService = &mockUserService{
					mockCreateUser: func(user *domain.SUser) (*domain.SUser, error) {
						return &domain.SUser{ID: "random id"}, nil
					},
				}
			},
			expectedCode: http.StatusCreated,
			expectedErr:  nil,
		},
	}
	for _, testCase := range testCases {
		testCase.setup()
		userPort := &SUserPort{
			UserService: userService,
		}

		route := SetUpRouter()
		route.POST("/api/user", userPort.CreateUser)

		jsonValue, _ := json.Marshal(testCase.req)
		req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}
