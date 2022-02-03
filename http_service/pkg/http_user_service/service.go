package httpuserservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"

	"github.com/jcromanu/final_project/http_service/errors"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
)

type HttpService struct {
	repo      Repository
	log       log.Logger
	validator *validator.Validate
}

type Repository interface {
	CreateUser(context.Context, entities.User) (int32, error)
	GetUser(context.Context, int32) (entities.User, error)
	UpdateUser(context.Context, entities.User) (string, error)
}

func NewHttpService(repo Repository, logger log.Logger) *HttpService {
	return &HttpService{
		repo:      repo,
		log:       logger,
		validator: validator.New(),
	}
}

func (srv *HttpService) CreateUser(ctx context.Context, usr entities.User) (entities.User, error) {
	if err := srv.validator.Struct(usr); err != nil {
		return entities.User{}, errors.NewEmptyFieldError()
	}
	id, err := srv.repo.CreateUser(ctx, usr)
	if err != nil {
		level.Error(srv.log).Log("Error creating user from grpc service  :", err)
		return entities.User{}, err
	}
	usr.Id = id
	return usr, err
}

func (srv *HttpService) GetUser(ctx context.Context, id int32) (entities.User, error) {
	if id <= 0 {
		return entities.User{}, errors.NewEmptyFieldError()
	}
	usr, err := srv.repo.GetUser(ctx, id)
	if err != nil {
		level.Error(srv.log).Log("Error getting user from grpc service  :", err)
		return entities.User{}, err
	}
	return usr, nil
}

func (srv *HttpService) UpdateUser(ctx context.Context, usr entities.User) (string, error) {
	if usr.Id <= 0 {
		return "", errors.NewEmptyFieldError()
	}
	res, err := srv.repo.UpdateUser(ctx, usr)
	if err != nil {
		level.Error(srv.log).Log("Error updating user from update service: ", err)
		return "", err
	}
	return res, nil
}
