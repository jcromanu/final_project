package userservice

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/pkg/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EndpointTestSuite struct {
	suite.Suite
	logger log.Logger
	ctx    context.Context
}

func (suite *EndpointTestSuite) SetupSuite() {
	suite.logger = log.NewLogfmtLogger(os.Stderr)
	suite.ctx = context.Background()
	level.Info(suite.logger).Log("Iniciando Endpoint test suite ")
}

func (suite *EndpointTestSuite) TestmakeCreateUserEndpoint() {
	serviceMock := new(ServiceMock)
	usr := entities.User{Id: 1}
	serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(usr, nil)
	ep := makeCreateUserEndpoint(serviceMock, suite.logger)
	req := createUserRequest{User: usr}
	result, err := ep(suite.ctx, req)
	if err != nil {
		suite.T().Errorf("Error creating user endpoint")
		return
	}
	re, ok := result.(createUserResponse)
	if !ok {
		suite.T().Errorf("Error parsing user response on test")
		return
	}
	assert.Equal(suite.T(), req.User.Id, re.User.Id, "Error on user request")
}

func TestEndpointTestSuite(t *testing.T) {
	suite.Run(t, new(EndpointTestSuite))
}
