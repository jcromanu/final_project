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
			testName:       "endpoint valid user",
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
