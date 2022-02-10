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
			Id:                    pbReq.User.Id,
			Name:                  pbReq.User.Name,
			PwdHash:               pbReq.User.PwdHash,
			Age:                   pbReq.User.Age,
			AdditionalInformation: pbReq.User.AdditionalInformation,
			Parent:                pbReq.User.Parent,
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
		return &pb.CreateUserResponse{User: &pb.User{Id: res.User.Id, PwdHash: res.User.PwdHash, Name: res.User.Name, Age: res.User.Age, Parent: res.User.Parent, AdditionalInformation: res.User.AdditionalInformation}, Message: &pb.MessageResponse{Code: res.Message.Code, Message: res.Message.Message}}, nil
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
		return &pb.GetUserResponse{User: &pb.User{Id: res.User.Id, PwdHash: res.User.PwdHash, Name: res.User.Name, Age: res.User.Age, AdditionalInformation: res.User.AdditionalInformation, Parent: res.User.Parent}, Message: &pb.MessageResponse{Code: res.Message.Code, Message: res.Message.Message}}, nil
	}
}

func makeDecodeGRPCUpdateUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.UpdateUserRequest)
		if !ok {
			level.Error(logger).Log("Update user request pb not matched")
			return nil, errors.NewProtoRequestError()
		}
		return updateUserRequest{entities.User{Id: pbReq.User.Id, PwdHash: pbReq.User.PwdHash, Name: pbReq.User.Name, Age: pbReq.User.Age, AdditionalInformation: pbReq.User.AdditionalInformation, Parent: pbReq.User.Parent}}, nil
	}
}

func makeEncodeGRPCUpdateUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(updateUserResponse)
		if !ok {
			level.Error(logger).Log("Get user response  pb not matched")
			return nil, errors.NewProtoResponseError()
		}
		return &pb.UpdateUserResponse{Message: &pb.MessageResponse{Code: res.Message.Code, Message: res.Message.Message}}, nil
	}
}

func makeDecodeDeleteUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, ok := req.(*pb.DeleteUserRequest)
		if !ok {
			level.Error(logger).Log("Get user response  pb not matched")
			return nil, errors.NewProtoResponseError()
		}
		return deleteUserRequest{id: res.Id}, nil
	}
}

func makeEncodeDeleteUserResponse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(deleteUserResponse)
		if !ok {
			level.Error(logger).Log("Get user response  pb not matched")
			return nil, errors.NewProtoResponseError()
		}
		return &pb.DeleteUserResponse{Message: &pb.MessageResponse{Message: res.Message.Message, Code: 0}}, nil
	}
}
