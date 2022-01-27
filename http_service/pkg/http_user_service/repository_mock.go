package httpuserservice

import (
	"context"

	"github.com/jcromanu/final_project/http_service/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type HttpRepositoryMock struct {
	mock.Mock
}

func (m *HttpRepositoryMock) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	args := m.Called(ctx, usr)
	return int32(args.Int(0)), args.Error(1)
}
