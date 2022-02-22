package userservice

import (
	"context"

	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pb"
)

func makeEncodeGRPCCReateUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(createUserResponse)
		if !ok {
			level.Error(logger).Log(errors.NewProtoResponseError().Error())
			return nil, errors.NewProtoResponseError()
		}
		return &pb.CreateUserResponse{
			User: &pb.User{
				Id:                    res.user.Id,
				PwdHash:               res.user.PwdHash,
				Name:                  res.user.Name,
				Age:                   res.user.Age,
				Parent:                res.user.Parent,
				AdditionalInformation: res.user.AdditionalInformation,
				Email:                 res.user.Email},
			Message: &pb.MessageResponse{
				Code:    res.message.Code,
				Message: res.message.Message}}, nil
	}
}

func makeEncodeGRPCGetUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(getUserResponse)
		if !ok {
			level.Error(logger).Log(errors.NewProtoResponseError().Error())
			return nil, errors.NewProtoResponseError()
		}
		return &pb.GetUserResponse{
			User: &pb.User{Id: res.user.Id,
				PwdHash:               res.user.PwdHash,
				Name:                  res.user.Name,
				Age:                   res.user.Age,
				AdditionalInformation: res.user.AdditionalInformation,
				Parent:                res.user.Parent,
				Email:                 res.user.Email},
			Message: &pb.MessageResponse{
				Code:    res.message.Code,
				Message: res.message.Message}}, nil
	}
}

func makeEncodeGRPCUpdateUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(updateUserResponse)
		if !ok {
			level.Error(logger).Log(errors.NewProtoResponseError().Error())
			return nil, errors.NewProtoResponseError()
		}
		return &pb.UpdateUserResponse{Message: &pb.MessageResponse{Code: res.message.Code, Message: res.message.Message}}, nil
	}
}

func makeEncodeDeleteUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(deleteUserResponse)
		if !ok {
			level.Error(logger).Log(errors.NewProtoResponseError())
			return nil, errors.NewProtoResponseError()
		}
		return &pb.DeleteUserResponse{Message: &pb.MessageResponse{Message: res.message.Message, Code: 0}}, nil
	}
}
