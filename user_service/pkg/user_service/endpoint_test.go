package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeCreateUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          createUserRequest
		expectedOutput createUserResponse
		expectedError  error
	}{
		{
			testName:       "test create endpoint user with all fields ",
			input:          createUserRequest{User: entities.User{Id: 0, Name: "Juan", PwdHash: "hash", Age: 30, AdditionalInformation: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: createUserResponse{User: entities.User{Id: 1}, Message: entities.Message{Message: "User created", Code: 0}},
			expectedError:  nil,
		},
		{
			testName:       "test create endpoint empty name ",
			input:          createUserRequest{User: entities.User{Id: 0, Name: "", PwdHash: "hash", Age: 30, AdditionalInformation: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: createUserResponse{},
			expectedError:  errors.NewBadRequestError(),
		},
	}
	for _, tc := range testCases {
		serviceMock := new(ServiceMock)
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput.User.Id, tc.expectedError)
			ep := makeCreateUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			assert.Equal(t, tc.expectedError, err, "Error on service create user  ")
			re, _ := result.(createUserResponse)
			assert.Equal(t, tc.expectedOutput, re, "Unexpected output")
		})
	}
}

func TestMakeGetUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          getUserRequest
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "test get endpoint user with all fields   ",
			input:          getUserRequest{Id: 1},
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock := new(ServiceMock)
			serviceMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.input.Id, tc.expectedError)
			ep := makeGetUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(getUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedOutput, re.User, "Error on user response")
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}

func TestMakeUpdateUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          updateUserRequest
		expectedOutput error
		expectedError  error
	}{
		{
			testName:       "test update endpoint user with all fields   ",
			input:          updateUserRequest{entities.User{Id: 1, Name: "Juan", Age: 30, PwdHash: "hash ", AdditionalInformation: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: nil,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock := new(ServiceMock)
			serviceMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			ep := makeUpdatesUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := result.(updateUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName      string
		input         deleteUserRequest
		expectedError error
	}{
		{
			testName:      "test delete user on request  ",
			input:         deleteUserRequest{id: 1},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock := new(ServiceMock)
			serviceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			ep := makeDeleteUserEndpoint(serviceMock, logger)
			res, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			_, ok := res.(deleteUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}

}
