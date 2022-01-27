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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransportTestSuite struct {
	suite.Suite
	logger      log.Logger
	ctx         context.Context
	middlewares []endpoint.Middleware
	opts        []grpc.ServerOption
}

func (suite *TransportTestSuite) SetupSuite() {
	suite.logger = log.NewLogfmtLogger(os.Stderr)
	suite.ctx = context.Background()
	suite.middlewares = []endpoint.Middleware{}
	suite.opts = []grpc.ServerOption{}
	level.Info(suite.logger).Log("Iniciando Transport test suite ")
}

func (suite *TransportTestSuite) TestCreateUser() {
	srvMock := new(ServiceMock)
	srvMock.On("CreateUser", mock.Anything, mock.Anything).Return(1, nil)
	endpoints := MakeEndpoints(srvMock, suite.logger, suite.middlewares)
	grpcServer := NewGRPCServer(endpoints, suite.opts, suite.logger)
	res, err := grpcServer.CreateUser(suite.ctx, &pb.CreateUserRequest{User: &pb.User{Id: 1}})
	if err != nil {
		suite.T().Errorf(err.Error())
		return
	}
	assert.Equal(suite.T(), int32(1), res.User.Id, "Different user id s")
}

func TestTransportTestSuite(t *testing.T) {
	suite.Run(t, new(TransportTestSuite))
}
