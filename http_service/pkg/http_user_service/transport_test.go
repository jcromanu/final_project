package httpuserservice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserTransport(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	middlewares := []endpoint.Middleware{}
	testCases := []struct {
		testName       string
		input          string
		expectedOutput int32
		expectedError  error
		expectedStatus int
	}{
		{
			testName: "test transport user created",
			input: `{"User":{"pwd_hash":"oooo",
					"name": "Juan ",
					"age": 30 ,
					"additional_information": "no info ", 
					"parent": ["testparent"]
					}}`,
			expectedOutput: 1,
			expectedError:  nil,
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)
			serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			endpoints := MakeEndpoints(serviceMock, logger, middlewares)
			httpServer := NewHTTPServer(endpoints, logger)
			server := httptest.NewServer(httpServer)
			res, _ := http.Post(server.URL+"/users", "application/json", strings.NewReader(tc.input))
			defer server.Close()
			assert.Equal(tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestGetUserTransport(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	middlewares := []endpoint.Middleware{}
	testCases := []struct {
		testName       string
		input          int32
		expectedOutput int32
		expectedError  error
		expectedStatus int
	}{
		{
			testName:       "test transport user retrieval ",
			input:          1,
			expectedOutput: 1,
			expectedError:  nil,
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)
			serviceMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			endpoints := MakeEndpoints(serviceMock, logger, middlewares)
			httpServer := NewHTTPServer(endpoints, logger)
			server := httptest.NewServer(httpServer)
			res, _ := http.Get(server.URL + "/users/" + fmt.Sprintf("%v", tc.expectedOutput))
			defer server.Close()
			assert.Equal(tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestUpdateUserTransport(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	middlewares := []endpoint.Middleware{}
	testCases := []struct {
		testName       string
		input          string
		id             int32
		expectedOutput int32
		expectedError  error
		expectedStatus int
	}{
		{
			testName: "test transport  user updated ",
			id:       1,
			input: `{"User":{"pwd_hash":"oooo",
			"name": "Juan ",
			"age": 30 ,
			"additional_information": "no info ", 
			"parent": ["testparent"]
			}}`,
			expectedOutput: 1,
			expectedError:  nil,
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)
			serviceMock.On("UpdateUser", mock.Anything, mock.Anything).Return("", tc.expectedError)
			endpoints := MakeEndpoints(serviceMock, logger, middlewares)
			httpServer := NewHTTPServer(endpoints, logger)
			server := httptest.NewServer(httpServer)
			req, _ := http.NewRequest("PUT", server.URL+"/users/"+fmt.Sprintf("%v", tc.id), strings.NewReader(tc.input))
			res, _ := http.DefaultClient.Do(req)
			defer server.Close()
			assert.Equal(tc.expectedStatus, res.StatusCode)
		})
	}
}

func TestDeleteUserTransport(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	middlewares := []endpoint.Middleware{}
	testCases := []struct {
		testName       string
		id             int32
		expectedOutput string
		expectedError  error
		expectedStatus int
	}{
		{
			testName:       "test transport user deleted ",
			id:             1,
			expectedOutput: "user deleted",
			expectedError:  nil,
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert := assert.New(t)
			serviceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			endpoints := MakeEndpoints(serviceMock, logger, middlewares)
			httpServer := NewHTTPServer(endpoints, logger)
			server := httptest.NewServer(httpServer)
			req, _ := http.NewRequest("DELETE", server.URL+"/users/"+fmt.Sprintf("%v", tc.id), nil)
			res, _ := http.DefaultClient.Do(req)
			defer server.Close()
			assert.Equal(tc.expectedStatus, res.StatusCode)
		})
	}
}
