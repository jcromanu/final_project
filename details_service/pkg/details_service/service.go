package userservice

//En esta capa nada relacionado con Gokit , ni transport , ni PB
//Crear entities que reflejen lo que se va a transportar

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

type UserService struct {
	repo      Repository
	logger    log.Logger
	validator Validator
}

type Repository interface {
	CreateUser(ctx context.Context, usr entities.User) (int32, error)
	GetUser(ctx context.Context, id int32) (entities.User, error)
	UpdateUser(ctx context.Context, usr entities.User) error
	DeleteUser(ctx context.Context, id int32) error
}

type Validator interface {
	Struct(s interface{}) error
}

type secret struct {
	hashSecret string `ENV:"HASH_SECRET"`
}

func NewService(repo Repository, logger log.Logger, validator Validator) *UserService {
	return &UserService{
		repo:      repo,
		logger:    logger,
		validator: validator,
	}
}

func (srv *UserService) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if err := srv.validator.Struct(user); err != nil {
		level.Error(srv.logger).Log(err.Error())
		return entities.User{}, errors.NewBadRequestError()
	}
	secret := secret{}
	if err := env.Parse(&secret); err != nil {
		level.Error(srv.logger).Log(err.Error())
		return entities.User{}, errors.NewInternalError()
	}
	checksum := sha256.Sum256([]byte(user.PwdHash + secret.hashSecret))
	user.PwdHash = string(fmt.Sprintf("%x", checksum))
	id, err := srv.repo.CreateUser(ctx, user)
	if err != nil {
		level.Error(srv.logger).Log(err.Error())
		return entities.User{}, err
	}
	user.Id = id
	user.PwdHash = ""
	return user, nil
}

func (srv *UserService) GetUser(ctx context.Context, id int32) (entities.User, error) {
	if id <= 0 {
		level.Error(srv.logger).Log(errors.NewBadRequestError().Error())
		return entities.User{}, errors.NewBadRequestError()
	}
	usr, err := srv.repo.GetUser(ctx, id)
	if err != nil {
		level.Error(srv.logger).Log(err.Error())
		return entities.User{}, err
	}
	return usr, nil
}

func (srv *UserService) UpdateUser(ctx context.Context, usr entities.User) error {
	if usr.Id <= 0 {
		level.Error(srv.logger).Log(errors.NewBadRequestError().Error())
		return errors.NewBadRequestError()
	}
	if err := srv.validator.Struct(usr); err != nil {
		level.Error(srv.logger).Log(err.Error())
		return errors.NewBadRequestError()
	}
	secret := secret{}
	if err := env.Parse(&secret); err != nil {
		level.Error(srv.logger).Log(err.Error())
		return errors.NewInternalError()
	}
	checksum := sha256.Sum256([]byte(usr.PwdHash + secret.hashSecret))
	usr.PwdHash = string(fmt.Sprintf("%x", checksum))
	err := srv.repo.UpdateUser(ctx, usr)
	if err != nil {
		level.Error(srv.logger).Log(err.Error())
		return err
	}
	return nil
}

func (srv *UserService) DeleteUser(ctx context.Context, id int32) error {
	if id <= 0 {
		level.Error(srv.logger).Log(errors.NewBadRequestError().Error())
		return errors.NewBadRequestError()
	}
	err := srv.repo.DeleteUser(ctx, id)
	if err != nil {
		level.Error(srv.logger).Log(err.Error())
		return err
	}
	return nil
}
