package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
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
			testName:       "test create user server  user with all fields success  ",
			input:          createUserRequest{User: entities.User{Name: "Juan", Age: 30, AdditionalInformation: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: 1,
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
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

func TestTransportGetUser(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	middlewares := []endpoint.Middleware{}
	opts := []grpc.ServerOption{}
	srvMock := new(ServiceMock)

	testCases := []struct {
		testName       string
		input          getUserRequest
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "test get user  server  user with all fields success  ",
			input:          getUserRequest{Id: 1},
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			srvMock.On("GetUser", mock.Anything, mock.Anything).Return(tc.input.Id, tc.expectedError)
			endpoints := MakeEndpoints(srvMock, logger, middlewares)
			grpcServer := NewGRPCServer(endpoints, opts, logger)
			res, err := grpcServer.GetUser(ctx, &pb.GetUserRequest{Id: tc.input.Id})
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			assert.Equal(t, tc.expectedOutput.Id, res.User.Id, "User not retrieved")
		})
	}
}

func TestTransportUpdateUser(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	middlewares := []endpoint.Middleware{}
	opts := []grpc.ServerOption{}
	srvMock := new(ServiceMock)

	testCases := []struct {
		testName       string
		input          *pb.UpdateUserRequest
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName:       "test update user server  user with all fields success  ",
			input:          &pb.UpdateUserRequest{User: &pb.User{Id: 1, Name: "Juan", Age: 30, PwdHash: "hash ", AdditionalInformation: "additional info", Parent: []string{"parent sample"}}},
			expectedOutput: entities.User{Id: 1},
			expectedError:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			srvMock.On("UpdateUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			endpoints := MakeEndpoints(srvMock, logger, middlewares)
			grpcServer := NewGRPCServer(endpoints, opts, logger)
			_, err := grpcServer.UpdateUser(ctx, tc.input)
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			assert.Equal(t, err, tc.expectedError, "User not retrieved")
		})
	}
}

func TestTransportDeleteUser(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	ctx := context.Background()
	middlewares := []endpoint.Middleware{}
	opts := []grpc.ServerOption{}
	srvMock := new(ServiceMock)

	testCases := []struct {
		testName      string
		input         *pb.DeleteUserRequest
		expectedError error
	}{
		{
			testName:      "test update user server  user with all fields success  ",
			input:         &pb.DeleteUserRequest{Id: 1},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			srvMock.On("DeleteUser", mock.Anything, mock.Anything).Return(tc.expectedError)
			endpoints := MakeEndpoints(srvMock, logger, middlewares)
			grpcServer := NewGRPCServer(endpoints, opts, logger)
			_, err := grpcServer.DeleteUser(ctx, tc.input)
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			assert.Equal(t, err, tc.expectedError, "User not retrieved")
		})
	}
}
