package httpuserservice

import (
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
			testName: "transport empty user creation",
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
