package httpuserservice

import (
	"context"

	"github.com/jcromanu/final_project/http_service/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type HTTPServiceMock struct {
	mock.Mock
}

func (m *HTTPServiceMock) CreateUser(ctx context.Context, usr entities.User) (entities.User, error) {
	args := m.Called(ctx, usr)
	usr.Id = int32(args.Int(0))
	return entities.User{Id: usr.Id}, args.Error(1)
}
