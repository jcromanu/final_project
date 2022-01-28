package httpuserservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-playground/validator/v10"

	"github.com/jcromanu/final_project/http_service/errors"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
)

type HTTPService interface {
	CreateUser(context.Context, entities.User) (entities.User, error)
}

type HttpService struct {
	repo      HttpRepository
	log       log.Logger
	validator *validator.Validate
}

func NewHttpService(repo HttpRepository, logger log.Logger) *HttpService {
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
