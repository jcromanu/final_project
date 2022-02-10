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

func TestServiceCreateUser(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          entities.User
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "test create user with all fields",
			input:          entities.User{Name: "Juan", Age: 30, Pwd_hash: "hash", Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			testName:       "test create user empty required field",
			input:          entities.User{Name: "Juan", Age: 30, Pwd_hash: "", Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedOutput: 0,
			expectedError:  errors.NewBadRequestError(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock := new(RepositoryMock)
			repoMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			service := NewService(repoMock, logger)
			usr, err := service.CreateUser(ctx, tc.input)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedOutput, usr.Id)
		})
	}
}

func TestServiceGetUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          int32
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "test get user valid id",
			input:          1,
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
		{
			testName:       "test get user empty id ",
			input:          0,
			expectedOutput: entities.User{},
			expectedError:  errors.NewBadRequestError(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.input, tc.expectedError)
			service := NewService(repoMock, logger)
			usr, err := service.GetUser(ctx, tc.input)
			assert.Equal(t, tc.expectedOutput, usr)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestServiceUpdateUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName      string
		input         entities.User
		expectedError error
	}{
		{
			testName:      "test update user valid id",
			input:         entities.User{Id: 1, Name: "Juan", Age: 30, Pwd_hash: "hash ", Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedError: nil,
		},
		{
			testName:      "test update user invalid id",
			input:         entities.User{Id: 0, Name: "Juan", Age: 30, Pwd_hash: "hash ", Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedError: errors.NewBadRequestError(),
		},
		{
			testName:      "test update user empty required field  ",
			input:         entities.User{Id: 1, Name: "", Age: 30, Pwd_hash: "hash ", Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedError: errors.NewBadRequestError(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			service := NewService(repoMock, logger)
			err := service.UpdateUser(ctx, tc.input)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestServiceDeleteUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName      string
		input         int32
		expectedError error
	}{
		{
			testName:      "test hard delete user valid id  ",
			input:         1,
			expectedError: nil,
		},
		{
			testName:      "test hard delete user invalid id",
			input:         0,
			expectedError: errors.NewBadRequestError(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			service := NewService(repoMock, logger)
			err := service.DeleteUser(ctx, tc.input)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
