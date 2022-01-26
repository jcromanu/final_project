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

type ServiceTestSuite struct {
	suite.Suite
	logger log.Logger
	ctx    context.Context
}

func (suite *ServiceTestSuite) SetupSuite() {
	suite.logger = log.NewLogfmtLogger(os.Stderr)
	suite.ctx = context.Background()
	level.Info(suite.logger).Log("Iniciando Service test suite ")
}

func (suite *ServiceTestSuite) TestCreateUserSuccess() {
	usr := entities.User{}
	repoMock := new(RepositoryMock)
	repoMock.On("CreateUser", mock.Anything, mock.Anything).Return(1, nil)
	service := NewService(repoMock, suite.logger)
	usr, actualErr := service.CreateUser(suite.ctx, usr)
	assert.Equal(suite.T(), int32(1), usr.Id)
	assert.Equal(suite.T(), nil, actualErr)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
