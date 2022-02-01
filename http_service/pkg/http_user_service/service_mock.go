package httpuserservice

import (
	"context"

	"github.com/jcromanu/final_project/http_service/pkg/entities"
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
	return entities.User{Id: id, Name: "Juan"}, args.Error(1)
}
