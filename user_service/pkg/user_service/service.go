package userservice

//En esta capa nada relacionado con Gokit , ni transport , ni PB
//Crear entities que reflejen lo que se va a transportar

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

type UserService struct {
	repo   Repository
	logger log.Logger
}

type Repository interface {
	CreateUser(context.Context, entities.User) (int32, error)
	GetUser(context.Context, int32) (entities.User, error)
	UpdateUser(context.Context, entities.User) error
	DeleteUser(context.Context, int32) error
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

func (srv *UserService) GetUser(ctx context.Context, id int32) (entities.User, error) {
	if id <= 0 {
		level.Error(srv.logger).Log("Empty user id ")
		return entities.User{}, errors.NewBadRequestError()
	}
	usr, err := srv.repo.GetUser(ctx, id)
	if err != nil {
		level.Error(srv.logger).Log("Error retrieving  user in database:", err)
		return entities.User{}, err
	}
	return usr, err
}

func (srv *UserService) UpdateUser(ctx context.Context, usr entities.User) error {
	if usr.Id <= 0 {
		level.Error(srv.logger).Log("Empty user id ")
		return errors.NewBadRequestError()
	}
	err := srv.repo.UpdateUser(ctx, usr)
	if err != nil {
		level.Error(srv.logger).Log("Error updating user in database:", err)
		return err
	}
	return nil
}

func (srv *UserService) DeleteUser(ctx context.Context, id int32) error {
	if id <= 0 {
		level.Error(srv.logger).Log("Empty user id ")
		return errors.NewBadRequestError()
	}
	err := srv.repo.DeleteUser(ctx, id)
	if err != nil {
		level.Error(srv.logger).Log("Error deleting user in database:", err)
		return err
	}
	return nil
}
