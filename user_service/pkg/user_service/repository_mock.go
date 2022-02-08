package userservice

import (
	"context"

	"github.com/jcromanu/final_project/user_service/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	args := m.Called(ctx, usr)
	id := args.Get(0).(int32)
	return id, args.Error(1)
}

func (m *RepositoryMock) GetUser(ctx context.Context, id int32) (entities.User, error) {
	args := m.Called(ctx, id)
	nId := args.Get(0).(int32)
	return entities.User{Id: nId}, args.Error(1)
}

func (m *RepositoryMock) UpdateUser(ctx context.Context, usr entities.User) error {
	args := m.Called(ctx, usr)
	return args.Error(0)
}

func (m *RepositoryMock) DeleteUser(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
