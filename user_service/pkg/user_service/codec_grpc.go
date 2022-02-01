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
		return createUserRequest{User: entities.User{
			Id:                     pbReq.User.Id,
			Name:                   pbReq.User.Name,
			Pwd_hash:               pbReq.User.PwdHash,
			Age:                    pbReq.User.Age,
			Additional_information: pbReq.User.AdditionalInformation,
			Parent:                 pbReq.User.Parent,
		}}, nil
	}
}

func makeEncodeGRPCCReateUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(createUserResponse)
		if !ok {
			level.Error(logger).Log("Create user response  pb not matched")
			return nil, errors.NewProtoResponseError()
		}
		return &pb.CreateUserResponse{User: &pb.User{Id: res.User.Id, PwdHash: res.User.Pwd_hash, Name: res.User.Name, Age: res.User.Age, Parent: res.User.Parent, AdditionalInformation: res.User.Additional_information}, Message: &pb.MessageResponse{Code: res.Message.Code, Message: res.Message.Message}}, nil
	}
}

func makeDecodeGRPCGetUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.GetUserRequest)
		if !ok {
			level.Error(logger).Log("Get user request pb not matched")
			return nil, errors.NewProtoRequestError()
		}
		return getUserRequest{pbReq.Id}, nil
	}
}

func makeEncodeGRPCGetUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(getUserResponse)
		if !ok {
			level.Error(logger).Log("Get user response  pb not matched")
			return nil, errors.NewProtoResponseError()
		}
		return &pb.GetUserResponse{User: &pb.User{Id: res.User.Id, PwdHash: res.User.Pwd_hash, Name: res.User.Name, Age: res.User.Age, AdditionalInformation: res.User.Additional_information}, Message: &pb.MessageResponse{Code: res.Message.Code, Message: res.Message.Message}}, nil
	}
}
