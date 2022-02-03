package httpuserservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeCreateHTTPUserEndpooint(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName       string
		input          createUserRequest
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "endpoint  create valid user",
			input:          createUserRequest{entities.User{Name: "Juan", Pwd_hash: "ooo", Age: 31, Additional_information: "no info", Parent: []string{}}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputRequest := tc.input
			serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			ep := makeCreateUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, inputRequest)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(createUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(tc.expectedOutput, re.User.Id, "Error on user request ")
		})
	}
}

func TestMakeGetHTTPUserEndpooint(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName       string
		input          getUserRequest
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "endpoint get valid user",
			input:          getUserRequest{Id: 1},
			expectedOutput: entities.User{Id: 1, Name: "Juan"},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputRequest := tc.input
			serviceMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.expectedOutput.Id, tc.expectedError)
			ep := makeGetUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, inputRequest)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(getUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(tc.expectedOutput, re.User, "Error on user request ")
		})
	}
}

func TestMakeUpdateHTTPUserEndpooint(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName        string
		input           updateUserRequest
		expectedError   error
		expectedMessage string
	}{
		{
			testName:        "endpoint get valid user",
			input:           updateUserRequest{entities.User{Id: 1, Name: "Juan", Pwd_hash: "ooo", Age: 31, Additional_information: "no info", Parent: []string{}}},
			expectedError:   nil,
			expectedMessage: "user updated",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputRequest := tc.input
			serviceMock.On("UpdateUser", mock.Anything, mock.Anything).Return("", tc.expectedError)
			ep := makeUpdateUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, inputRequest)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := result.(updateUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(tc.expectedError, err, "Error on user request ")
		})
	}
}
