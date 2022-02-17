package httpuserservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/jcromanu/final_project/http_service/errors"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
)

type HttpService struct {
	repo      Repository
	log       log.Logger
	validator Validator
}

type Repository interface {
	CreateUser(ctx context.Context, usr entities.User) (int32, error)
	GetUser(ctx context.Context, id int32) (entities.User, error)
	UpdateUser(ctx context.Context, usr entities.User) (string, error)
	DeleteUser(ctx context.Context, id int32) (string, error)
}

type Validator interface {
	Struct(s interface{}) error
}

func NewHttpService(repo Repository, logger log.Logger, validator Validator) *HttpService {
	return &HttpService{
		repo:      repo,
		log:       logger,
		validator: validator,
	}
}

func (srv *HttpService) CreateUser(ctx context.Context, usr entities.User) (entities.User, error) {
	if err := srv.validator.Struct(usr); err != nil {
		level.Error(srv.log).Log(err.Error())
		return entities.User{}, errors.NewEmptyFieldError()
	}
	id, err := srv.repo.CreateUser(ctx, usr)
	if err != nil {
		level.Error(srv.log).Log(err.Error())
		return entities.User{}, err
	}
	usr.Id = id
	usr.PwdHash = ""
	return usr, nil
}

func (srv *HttpService) GetUser(ctx context.Context, id int32) (entities.User, error) {
	if id <= 0 {
		level.Error(srv.log).Log(errors.NewEmptyFieldError().Error())
		return entities.User{}, errors.NewEmptyFieldError()
	}
	usr, err := srv.repo.GetUser(ctx, id)
	if err != nil {
		level.Error(srv.log).Log(err.Error())
		return entities.User{}, err
	}
	return usr, nil
}

func (srv *HttpService) UpdateUser(ctx context.Context, usr entities.User) (string, error) {
	if usr.Id <= 0 {
		level.Error(srv.log).Log(errors.NewEmptyFieldError().Error())
		return "", errors.NewEmptyFieldError()
	}
	if err := srv.validator.Struct(usr); err != nil {
		level.Error(srv.log).Log(err.Error())
		return "", errors.NewEmptyFieldError()
	}

	res, err := srv.repo.UpdateUser(ctx, usr)
	if err != nil {
		level.Error(srv.log).Log(err.Error())
		return "", err
	}
	return res, nil
}

func (srv *HttpService) DeleteUser(ctx context.Context, id int32) (string, error) {
	if id <= 0 {
		level.Error(srv.log).Log(errors.NewEmptyFieldError().Error())
		return "", errors.NewEmptyFieldError()
	}
	res, err := srv.repo.DeleteUser(ctx, id)
	if err != nil {
		level.Error(srv.log).Log(err.Error())
		return "", err
	}
	return res, nil
}
