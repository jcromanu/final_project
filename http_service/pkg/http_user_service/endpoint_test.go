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
			testName:       "test create user endpoint ",
			input:          createUserRequest{Name: "Juan", PwdHash: "ooo", Age: 31, AdditionalInformation: "no info", Parent: []string{}},
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
			assert.Equal(tc.expectedOutput, re.Id, "Error on user request ")
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
			testName:       "test get user endpoint",
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
			assert.Equal(tc.expectedOutput, entities.User{Id: re.Id, Name: re.Name}, "Error on user request ")
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
			testName:        "test update user endpoint",
			input:           updateUserRequest{Id: 1, Name: "Juan", PwdHash: "ooo", Age: 31, AdditionalInformation: "no info", Parent: []string{}},
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

func TestMakeDeleteHTTPUserEndpooint(t *testing.T) {
	serviceMock := new(ServiceMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName        string
		input           deleteUserRequest
		expectedError   error
		expectedMessage string
	}{
		{
			testName:        "test delete user endpoint",
			input:           deleteUserRequest{Id: 1},
			expectedError:   nil,
			expectedMessage: "user deleted",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputRequest := tc.input
			serviceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedMessage, tc.expectedError)
			ep := makeDeleteUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, inputRequest)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := result.(deleteUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(tc.expectedError, err, "Error on user request ")
		})
	}
}
