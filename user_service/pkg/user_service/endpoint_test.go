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

func TestMakeCreateUserEndpoint(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	serviceMock := new(ServiceMock)
	testCases := []struct {
		testName       string
		input          createUserRequest
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "create endpoint user with all fields success ",
			input:          createUserRequest{User: entities.User{Name: "Juan", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			ep := makeCreateUserEndpoint(serviceMock, logger)
			result, err := ep(ctx, tc.input)
			if err != nil {
				t.Errorf("Error creating user endpoint")
				return
			}
			re, ok := result.(createUserResponse)
			if !ok {
				t.Errorf("Error parsing user response on test")
				return
			}
			assert.Equal(t, tc.expectedOutput, re.User.Id, "Error on user response")
			assert.Equal(t, tc.expectedError, err, "Error on user response")
		})
	}
}
