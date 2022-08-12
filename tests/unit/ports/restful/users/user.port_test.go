package portsRestFulUser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

// func (mock *MockRepository) FindAll() ([]entity.Post, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.([]entity.Post), args.Error(1)
// }

// func TestFindAll(t *testing.T) {
// 	mockRepo := new(MockRepository)

// 	var identifier int64 = 1

// 	post := entity.Post{ID: identifier, Title: "A", Text: "B"}
// 	// Setup expectations
// 	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

// 	testService := NewPostService(mockRepo)

// 	result, _ := testService.FindAll()

// 	//Mock Assertion: Behavioral
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertion
// 	assert.Equal(t, identifier, result[0].ID)
// 	assert.Equal(t, "A", result[0].Title)
// 	assert.Equal(t, "B", result[0].Text)
// }

func (mock *MockRepository) CreateUser(user *domain.SUser) (*domain.SUser, error) {
	args := mock.Called()
	result := args.Get(0)
	fmt.Println("\n\n\n\n testing createUser ")
	// fmt.Println("\n\n\n Test Create User \n\n")
	return result.(*domain.SUser), args.Error(1)
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

func TestCreateUser(t *testing.T) {
	t.Parallel()
	response := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(response)
	c.Request = &http.Request{
		Header: make(http.Header),
	}
	//============================
	// value := []byte(`{id: "123", Username: "testing username", LimitPerDay: 10}`)

	// response.WriteString(string(value))
	// fmt.Println("\n\n\n\n\n\n\nin the test case", response.Result().StatusCode)
	// byteData := ioutil.NopCloser(bytes.NewBuffer(value))
	// c.Request.Body = (byteData)

	mockRepo := new(MockRepository)
	mockPort := portsRestFulUser.NewUserPort(&pg.DB{})

	now := time.Now()
	user := domain.SUser{ID: "123", Username: "testing username", LimitPerDay: 10, UpdatedAt: now, CreatedAt: now}

	mockRepo.On("CreateUser").Return(&user, nil)

	// MockJsonPost(c, map[string]interface{}{"id": "123", "Username": "testing username", "LimitPerDay": 10})
	// Run function
	mockPort.CreateUser(c)

	// fmt.Println("\n\n\n\n testing ", c)

	// assert.Equal(t, 400, response.Result().StatusCode)
	assert.Equal(t, 300, response.Code)

}
