package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"

	"github.com/jcromanu/final_project/user_service/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceCreateUser(t *testing.T) {
	repoMock := new(RepositoryMock)
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	testCases := []struct {
		testName       string
		input          entities.User
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "create user with all fields",
			input:          entities.User{Name: "Juan", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}},
			expectedOutput: int32(1),
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoMock.On("CreateUser", mock.Anything, mock.Anything).Return(1, nil)
			service := NewService(repoMock, logger)
			usr, err := service.CreateUser(ctx, tc.input)
			assert.Equal(t, tc.expectedOutput, usr.Id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
