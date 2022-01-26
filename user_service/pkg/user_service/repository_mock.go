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
	return int32(args.Int(0)), args.Error(1)
}
