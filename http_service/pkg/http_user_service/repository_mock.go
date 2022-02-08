package httpuserservice

import (
	"context"

	"github.com/jcromanu/final_project/http_service/pkg/entities"
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
	return entities.User{Id: id, Name: "Juan"}, args.Error(1)
}

func (m *RepositoryMock) UpdateUser(ctx context.Context, usr entities.User) (string, error) {
	args := m.Called(ctx, usr)
	return args.String(0), args.Error(1)
}

func (m *RepositoryMock) DeleteUser(ctx context.Context, id int32) (string, error) {
	args := m.Called(ctx, id)
	return args.String(0), args.Error(1)
}
