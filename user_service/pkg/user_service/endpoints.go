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
	CreateUser endpoint.Endpoint
}

func MakeEndpoints(srv Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		//CreateUser: wrapEndpoints(makeCreateUserEndpoint(srv, logger), middlewares), completar cuando vea lo de middlewares
		CreateUser: makeCreateUserEndpoint(srv, logger),
	}
}

func makeCreateUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createUserRequest)
		if !ok {
			level.Error(logger).Log("Bad request on endpoint creation  expected createUserRequest got :")
			level.Error(logger).Log(reflect.TypeOf(request))
			return nil, errors.NewBadRequestError()
		}
		usr, err := srv.CreateUser(ctx, req.User)
		if err != nil {
			level.Error(logger).Log(err)
			return nil, err
		}
		return createUserResponse{User: usr, Message: entities.Message{Message: "User created", Code: 0}}, nil
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
