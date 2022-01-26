package userservice

import (
	"context"

	"github.com/jcromanu/final_project/user_service/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) CreateUser(ctx context.Context, usr entities.User) (entities.User, error) {
	args := m.Called(ctx, usr)
	return entities.User{Id: usr.Id}, args.Error(1)
}
