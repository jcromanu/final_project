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
	usr.Id = args.Get(0).(int32)
	return entities.User{Id: usr.Id}, args.Error(1)
}

func (m *ServiceMock) GetUser(ctx context.Context, id int32) (entities.User, error) {
	args := m.Called(ctx, id)
	parsedID := args.Get(0).(int32)
	usr := entities.User{Id: parsedID}
	return usr, args.Error(1)
}

func (m *ServiceMock) UpdateUser(ctx context.Context, usr entities.User) error {
	args := m.Called(ctx, usr)
	return args.Error(0)
}
