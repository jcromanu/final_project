package userservice

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

type Endpoints struct {
	createUser endpoint.Endpoint
	getUser    endpoint.Endpoint
	updateUser endpoint.Endpoint
	deleteUser endpoint.Endpoint
}

type Service interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
	GetUser(ctx context.Context, id int32) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, id int32) error
}

func MakeEndpoints(srv Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		//CreateUser: wrapEndpoints(makeCreateUserEndpoint(srv, logger), middlewares)
		createUser: makeCreateUserEndpoint(srv, logger),
		getUser:    makeGetUserEndpoint(srv, logger),
		updateUser: makeUpdatesUserEndpoint(srv, logger),
		deleteUser: makeDeleteUserEndpoint(srv, logger),
	}
}

func makeCreateUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return createUserResponse{}, errors.NewBadRequestError()
		}
		usr, err := srv.CreateUser(ctx, req.user)
		if err != nil {
			level.Error(logger).Log(err)
			return createUserResponse{}, err
		}
		return createUserResponse{user: usr, message: entities.Message{Message: "User created", Code: 0}}, nil
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
		usr, err := srv.GetUser(ctx, req.id)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return getUserResponse{user: usr, message: entities.Message{Message: "User retrieved", Code: 0}}, nil
	}
}

func makeUpdatesUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewBadRequestError().Error())
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		err := srv.UpdateUser(ctx, req.user)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return updateUserResponse{entities.Message{Message: "user updated", Code: 0}}, nil
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
		err := srv.DeleteUser(ctx, req.id)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return deleteUserResponse{entities.Message{Message: "user deleted", Code: 0}}, nil
	}
}

/*
func wrapEndpoints(ep endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, middleware := range middlewares {
		ep = middleware(ep)
	}
	return ep
}

func middleware(ep endpoint.Middleware) endpoint.Endpoint {
	return nil
}
*/
