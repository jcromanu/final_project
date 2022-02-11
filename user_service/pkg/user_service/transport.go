package userservice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pb"
)

type userServiceServer struct {
	createUser *kitGRPC.Server
	getUser    *kitGRPC.Server
	updateUser *kitGRPC.Server
	deleteUser *kitGRPC.Server
	pb.UnimplementedUserServiceServer
}

func makeCreateUserGRPCServer(ep endpoint.Endpoint, opts []kitGRPC.ServerOption, logger log.Logger) *kitGRPC.Server {
	return kitGRPC.NewServer(
		ep,
		makeDecodeGRPCCreateUserRequest(logger),
		makeEncodeGRPCCReateUserResponse(logger),
		opts...,
	)
}

func makeGetUserGRPCServer(ep endpoint.Endpoint, opts []kitGRPC.ServerOption, logger log.Logger) *kitGRPC.Server {
	return kitGRPC.NewServer(
		ep,
		makeDecodeGRPCGetUserRequest(logger),
		makeEncodeGRPCGetUserResponse(logger),
		opts...,
	)
}

func makeUpdateUserGRPCServer(ep endpoint.Endpoint, opts []kitGRPC.ServerOption, logger log.Logger) *kitGRPC.Server {
	return kitGRPC.NewServer(
		ep,
		makeDecodeGRPCUpdateUserRequest(logger),
		makeEncodeGRPCUpdateUserResponse(logger),
		opts...,
	)
}

func makeDeleteUserGRPCServer(ep endpoint.Endpoint, opts []kitGRPC.ServerOption, logger log.Logger) *kitGRPC.Server {
	return kitGRPC.NewServer(
		ep,
		makeDecodeDeleteUserRequest(logger),
		makeEncodeDeleteUserResponse(logger),
		opts...,
	)
}

func NewGRPCServer(ep Endpoints, opts []kitGRPC.ServerOption, log log.Logger) pb.UserServiceServer {
	return &userServiceServer{
		createUser: makeCreateUserGRPCServer(ep.createUser, opts, log),
		getUser:    makeGetUserGRPCServer(ep.getUser, opts, log),
		updateUser: makeUpdateUserGRPCServer(ep.updateUser, opts, log),
		deleteUser: makeDeleteUserGRPCServer(ep.deleteUser, opts, log),
	}
}

func (srv *userServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := srv.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := resp.(*pb.CreateUserResponse)
	if !ok {
		return nil, errors.NewBadResponseTypeError()
	}
	return r, nil
}

func (srv *userServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := srv.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := resp.(*pb.GetUserResponse)
	if !ok {
		return nil, errors.NewBadResponseTypeError()
	}
	return r, nil
}

func (srv *userServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, resp, err := srv.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := resp.(*pb.UpdateUserResponse)
	if !ok {
		return nil, errors.NewBadResponseTypeError()
	}
	return r, nil
}

func (srv *userServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, resp, err := srv.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r, ok := resp.(*pb.DeleteUserResponse)
	if !ok {
		return nil, errors.NewBadResponseTypeError()
	}
	return r, nil
}
