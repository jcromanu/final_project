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

func TestCreateUser(t *testing.T) {
	httpRepoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName       string
		input          entities.User
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "test create user all fields   ",
			input:          entities.User{Name: "Juan", Pwd_hash: "ooo", Age: 31, Additional_information: "no info", Parent: []string{}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputUser := tc.input
			httpRepoMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			httpSrv := NewHttpService(httpRepoMock, logger)
			usr, err := httpSrv.CreateUser(ctx, inputUser)
			assert.Equal(tc.expectedOutput, usr.Id, "User creation fail ")
			assert.Equal(tc.expectedError, err, "User creation error:  ", err)
		})
	}
}

func TestGetUser(t *testing.T) {
	httpRepoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName       string
		input          int32
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "test get user with id   ",
			input:          1,
			expectedOutput: entities.User{Id: 1, Name: "Juan"},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			inputId := tc.input
			httpRepoMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			httpSrv := NewHttpService(httpRepoMock, logger)
			usr, err := httpSrv.GetUser(ctx, inputId)
			assert.Equal(tc.expectedOutput, usr, "User retrieval  fail ")
			assert.Equal(tc.expectedError, err, "User retrieval error:  ", err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	httpRepoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName        string
		input           entities.User
		expectedError   error
		expectedMessage string
	}{
		{
			testName:        "test update user all fields   ",
			input:           entities.User{Id: 1, Name: "Juan", Pwd_hash: "hash", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedError:   nil,
			expectedMessage: "user updated",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := context.Background()
			assert := assert.New(t)
			httpRepoMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedMessage, tc.expectedError)
			httpSrv := NewHttpService(httpRepoMock, logger)
			_, err := httpSrv.UpdateUser(ctx, tc.input)
			assert.Equal(tc.expectedError, err, "User retrieval fail ")
		})
	}
}

func DeleteUser(t *testing.T) {
	httpRepoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	testCases := []struct {
		testName        string
		input           int32
		expectedError   error
		expectedMessage string
	}{
		{
			testName:        "test delete user with id  ",
			input:           1,
			expectedError:   nil,
			expectedMessage: "user deleted",
		},
	}
	for _, tc := range testCases {
		ctx := context.Background()
		assert := assert.New(t)
		httpRepoMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedMessage, tc.expectedError)
		httpSrv := NewHttpService(httpRepoMock, logger)
		_, err := httpSrv.DeleteUser(ctx, tc.input)
		assert.Equal(tc.expectedError, err)
	}
}
