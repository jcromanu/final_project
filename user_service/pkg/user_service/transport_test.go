package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/pb"
	"github.com/jcromanu/final_project/user_service/pkg/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransportCreateUser(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	middlewares := []endpoint.Middleware{}
	opts := []grpc.ServerOption{}
	srvMock := new(ServiceMock)

	testCases := []struct {
		testName       string
		input          createUserRequest
		expectedOutput int32
		expectedError  error
	}{
		{
			testName:       "test serve endpoint user with all fields success  ",
			input:          createUserRequest{User: entities.User{Name: "Juan", Age: 30, Additional_information: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			level.Info(logger).Log("Iniciando Transport test suite ")
			srvMock.On("CreateUser", mock.Anything, mock.Anything).Return(tc.expectedOutput, tc.expectedError)
			endpoints := MakeEndpoints(srvMock, logger, middlewares)
			grpcServer := NewGRPCServer(endpoints, opts, logger)
			res, err := grpcServer.CreateUser(ctx, &pb.CreateUserRequest{User: &pb.User{Id: 1}})
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			assert.Equal(t, tc.expectedOutput, res.User.Id, "Different user id s")
		})
	}
}
