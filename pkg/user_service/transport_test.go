package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/pb"
	"github.com/jcromanu/final_project/pkg/entities"

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

/* Segunda opcion de implementacion de pruebas

type serverGRPCMock struct {
	mock.Mock
}

func (m *serverGRPCMock) ServeGRPC(ctx context.Context, req interface{}) (rect context.Context, resp interface{}, err error) {
	args := m.Called(ctx, req)
	return args.Get(0).(context.Context), args.Get(1), nil
}

func (suite *TransportTestSuite) TestCreateUser() {
	serverGRPCMock := &serverGRPCMock{}
	pbMockRes := &pb.CreateUserResponse{User: &pb.User{Id: 1}}
	serverGRPCMock.On("ServeGRPC", mock.Anything, mock.Anything).Return(suite.ctx, pbMockRes, nil)
	endpoints := Endpoints{}
	opts := []kitGRPC.ServerOption{}
	grpcServer := NewGRPCServer(endpoints, opts, suite.logger)
	usr := &pb.User{Id: 1}
	req := &pb.CreateUserRequest{User: usr}
	finalServer := grpcServer.(*userServiceServer)
	finalServer.createUser :=
	res, err := finalServer.CreateUser(suite.ctx, req)
	assert.Equal(suite.T(), req.User.Id, res.User.Id, "Error on assertion")
	assert.Equal(suite.T(), nil, err, "Error on assertion")
}*/

func (suite *TransportTestSuite) TestCreateUser() {
	srvMock := new(ServiceMock)
	usrReq := &entities.User{Id: 1}
	srvMock.On("CreateUser", mock.Anything, mock.Anything).Return(usrReq, nil)
	endpoints := MakeEndpoints(srvMock, suite.logger, suite.middlewares)
	grpcServer := NewGRPCServer(endpoints, suite.opts, suite.logger)
	res, err := grpcServer.CreateUser(suite.ctx, &pb.CreateUserRequest{User: &pb.User{Id: 1}})
	if err != nil {
		suite.T().Errorf(err.Error())
		return
	}
	assert.Equal(suite.T(), usrReq.Id, res.User.Id, "Different user id s")
}

func TestTransportTestSuite(t *testing.T) {
	suite.Run(t, new(TransportTestSuite))
}
