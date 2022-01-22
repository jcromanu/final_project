package userservice

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"github.com/jcromanu/final_project/pb"
	"github.com/jcromanu/final_project/pkg/entities"
)

var (
	errInternalError = errors.New("Internal error parsing request ")
)

func makeDecodeGRPCCreateUserRequest(logger log.Logger) kitGRPC.DecodeRequestFunc {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.CreateUserRequest)
		if !ok {
			return nil, errInternalError
		}
		return &createUserRequest{User: entities.User{
			Id:                     pbReq.User.Id,
			Pwd_hash:               pbReq.User.PwdHash,
			Age:                    pbReq.User.Age,
			Additional_information: pbReq.User.AdditionalInformation,
			Parent:                 pbReq.User.Parent,
		}}, nil
	}
}

func makeEncodeGRPCCReateUserResonse(logger log.Logger) kitGRPC.EncodeResponseFunc {
	return func(ctx context.Context, resp interface{}) (request interface{}, err error) {
		res, ok := resp.(createUserResponse)
		if !ok {
			return nil, errInternalError
		}
		return pb.CreateUserResponse{User: &pb.User{Id: res.user.Id, PwdHash: res.user.Pwd_hash, Name: res.user.Name, Age: res.user.Age, Parent: res.user.Parent}, Message: &pb.MessageResponse{Code: res.message.Code, Message: res.message.Message}}, nil
	}
}
