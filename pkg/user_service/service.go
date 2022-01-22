package userservice

//En esta capa nada relacionado con Gokit , ni transport , ni PB
//Crear entities que reflejen lo que se va a transportar

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jcromanu/final_project/pkg/entities"
)

var (
	errorUserCreationFail = errors.New("user creation failed ")
)

type Service interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
}

type UserService struct {
	repo   UserRepository
	logger log.Logger
}

func NewService(repo UserRepository, logger log.Logger) *UserService {
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
