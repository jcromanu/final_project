package userservice

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/jcromanu/final_project/pkg/entities"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
}

var (
	errorBadRequest = errors.New("Bad request for method ")
)

func MakeEndpoints(srv Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		//CreateUser: wrapEndpoints(makeCreateUserEndpoint(srv, logger), middlewares), completar cuando vea lo de middlewares
		CreateUser: makeCreateUserEndpoint(srv, logger),
	}
}

func makeCreateUserEndpoint(srv Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//logger := getLogger("makeCreateUserEndpoint",request_id.ReqID(ctx),logger)
		req, ok := request.(createUserRequest)
		if !ok {
			return nil, errorBadRequest
		}
		usr, err := srv.CreateUser(ctx, req.User)
		if err != nil {
			return createUserResponse{user: usr, message: entities.Message{Message: "User created", Code: 0}}, nil
		}
		return nil, err
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
