package userservice

import (
	"context"

	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/user_service/errors"
	"github.com/jcromanu/final_project/user_service/pb"
	"github.com/jcromanu/final_project/user_service/pkg/entities"
)

func makeDecodeGRPCCreateUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.CreateUserRequest)
		if !ok {
			level.Error(logger).Log("Create user request pb not matched")
			return nil, errors.NewProtoRequestError()
		}
		return createUserRequest{user: entities.User{
			Id:                    pbReq.User.Id,
			Name:                  pbReq.User.Name,
			PwdHash:               pbReq.User.PwdHash,
			Age:                   pbReq.User.Age,
			AdditionalInformation: pbReq.User.AdditionalInformation,
			Parent:                pbReq.User.Parent,
			Email:                 pbReq.User.Email,
		}}, nil
	}
}

func makeDecodeGRPCGetUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.GetUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewProtoRequestError().Error())
			return nil, errors.NewProtoRequestError()
		}
		return getUserRequest{pbReq.Id}, nil
	}
}

func makeDecodeGRPCUpdateUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.UpdateUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewProtoRequestError().Error())
			return nil, errors.NewProtoRequestError()
		}
		return updateUserRequest{entities.User{
				Id:                    pbReq.User.Id,
				PwdHash:               pbReq.User.PwdHash,
				Name:                  pbReq.User.Name,
				Age:                   pbReq.User.Age,
				AdditionalInformation: pbReq.User.AdditionalInformation,
				Parent:                pbReq.User.Parent,
				Email:                 pbReq.User.Email}},
			nil
	}
}

func makeDecodeDeleteUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, ok := req.(*pb.DeleteUserRequest)
		if !ok {
			level.Error(logger).Log(errors.NewProtoResponseError().Error())
			return nil, errors.NewProtoResponseError()
		}
		return deleteUserRequest{id: res.Id}, nil
	}
}
