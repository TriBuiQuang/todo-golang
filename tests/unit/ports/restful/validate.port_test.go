package portsRestFul

import (
	"net/http/httptest"
	"testing"
	"togo/internal/core/domain"
	portsRestFul "togo/internal/ports/restful"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func TestValidateCreateTask(t *testing.T) {
	t.Parallel()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// task *domain.STask

	// result :=

	// assert.Equal(t, true, result)

	testCases := []struct {
		name      string
		expect    bool
		testValue *domain.STask
	}{
		{name: "empty task", expect: true, testValue: &domain.STask{}},
		{name: "only have name", expect: true, testValue: &domain.STask{Name: "testing"}},
		{name: "only have userId", expect: true, testValue: &domain.STask{UserID: "123"}},
		{name: "have both name and userId but empty string", expect: true, testValue: &domain.STask{Name: "", UserID: ""}},
		{name: "normal case have both name and userId", expect: false, testValue: &domain.STask{Name: "testing", UserID: "123"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expect, portsRestFul.ValidateCreateTask(c, testCase.testValue))
		})
	}
}
