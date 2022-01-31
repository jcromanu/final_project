package userservice

//En esta capa nada relacionado con Gokit , ni transport , ni PB
//Crear entities que reflejen lo que se va a transportar

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

type UserService struct {
	repo   Repository
	logger log.Logger
}

type Repository interface {
	CreateUser(context.Context, entities.User) (int32, error)
}

func NewService(repo Repository, logger log.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (srv *UserService) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	id, err := srv.repo.CreateUser(ctx, user)
	if err != nil {
		level.Error(srv.logger).Log("Error creating user in database:", err)
		return entities.User{}, err
	}
	user.Id = id
	return user, err
}
