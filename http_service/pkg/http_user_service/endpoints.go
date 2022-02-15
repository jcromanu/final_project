package httpuserservice

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/http_service/errors"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
)

type Endpoints struct {
	createUser endpoint.Endpoint
	getUser    endpoint.Endpoint
	updateUser endpoint.Endpoint
	deleteUser endpoint.Endpoint
}

type Service interface {
	CreateUser(ctx context.Context, usr entities.User) (entities.User, error)
	GetUser(ctx context.Context, id int32) (entities.User, error)
	UpdateUser(ctx context.Context, usr entities.User) (string, error)
	DeleteUser(ctx context.Context, id int32) (string, error)
}

func MakeEndpoints(srv Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		createUser: makeCreateUserEndpoint(srv, logger),
		getUser:    makeGetUserEndpoint(srv, logger),
		updateUser: makeUpdateUserEndpoint(srv, logger),
		deleteUser: makeDeleteUserEndpoint(srv, logger),
	}
}

func makeCreateUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		usr, err := srv.CreateUser(ctx, entities.User{Name: req.Name, PwdHash: req.PwdHash, Age: req.Age, AdditionalInformation: req.AdditionalInformation, Parent: req.Parent, Email: req.Email})
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return createUserResponse{Id: usr.Id, PwdHash: usr.PwdHash, Name: usr.Name, Age: usr.Age, AdditionalInformation: usr.AdditionalInformation, Parent: usr.Parent, Email: usr.Email}, nil
	}
}

func makeGetUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		usr, err := srv.GetUser(ctx, req.Id)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return getUserResponse{Id: usr.Id, PwdHash: usr.PwdHash, Name: usr.Name, Age: usr.Age, AdditionalInformation: usr.AdditionalInformation, Parent: usr.Parent, Email: usr.Email}, nil
	}
}

func makeUpdateUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		res, err := srv.UpdateUser(ctx, entities.User{Id: req.Id, PwdHash: req.PwdHash, Name: req.Name, Age: req.Age, AdditionalInformation: req.AdditionalInformation, Parent: req.Parent, Email: req.Email})
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return updateUserResponse{Status: res}, nil
	}
}

func makeDeleteUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		res, err := srv.DeleteUser(ctx, req.Id)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return deleteUserResponse{Status: res}, nil
	}
}
