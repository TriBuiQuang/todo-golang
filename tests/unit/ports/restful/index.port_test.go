package portsRestFul

import (
	"errors"
	"net/http"
	"testing"
	mockRestFul "togo/__mocks__/ports/restful"
	portsRestFul "togo/internal/ports/restful"

	"github.com/stretchr/testify/assert"
)

func TestPrintErrResponse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		expectHttpStatus int
		expectErrMess    map[string]any
		testErrVal       error
		testHttpStatus   int
	}{
		{
			name:             "Information case",
			expectHttpStatus: 100,
			expectErrMess:    map[string]any{"message": "Information - testing", "status": 100},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   100,
		},
		{
			name:             "Success case",
			expectHttpStatus: 200,
			expectErrMess:    map[string]any{"message": "Success - testing", "status": 200},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   200,
		},
		{
			name:             "Redirection case",
			expectHttpStatus: 300,
			expectErrMess:    map[string]any{"message": "Redirection - testing", "status": 300},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   300,
		},
		{
			name:             "Client error case",
			expectHttpStatus: 400,
			expectErrMess:    map[string]any{"message": "Client error - testing", "status": 400},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   400,
		},
		{
			name:             "Server Error case",
			expectHttpStatus: 500,
			expectErrMess:    map[string]any{"message": "Server Error - testing", "status": 500},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   500,
		},
		{
			name:             "default",
			expectHttpStatus: 600,
			expectErrMess:    map[string]any{"message": "Don't have this status code- everything good - testing", "status": 600},
			testErrVal:       errors.New("testing"),
			testHttpStatus:   600,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			httpStatus, errMessage := portsRestFul.PrintErrResponse(testCase.testErrVal, testCase.testHttpStatus)
			assert.Equal(t, testCase.expectHttpStatus, httpStatus)
			assert.Equal(t, testCase.expectErrMess, errMessage)
		})
	}
}

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	response, c := mockRestFul.MockSetup()

	testCases := []struct {
		name   string
		actual int
	}{
		{name: "normal case", actual: 200},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			portsRestFul.HealthCheck(c)

			assert.Equal(t, response.Code, testCase.actual)
		})
	}
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	response, c := mockRestFul.MockSetup()

	testCases := []struct {
		name   string
		actual int
	}{
		{name: "normal case", actual: http.StatusNotFound},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			portsRestFul.NotFound(c)

			assert.Equal(t, response.Code, testCase.actual)
		})
	}
}

func TestWelcome(t *testing.T) {
	t.Parallel()

	response, c := mockRestFul.MockSetup()

	testCases := []struct {
		name   string
		actual int
	}{
		{name: "normal case", actual: http.StatusOK},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			portsRestFul.Welcome(c)

			assert.Equal(t, response.Code, testCase.actual)
		})
	}
}
