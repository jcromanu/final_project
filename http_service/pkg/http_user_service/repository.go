package httpuserservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/http_service/errors"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
	"github.com/jcromanu/final_project/user_service/pb"
	"google.golang.org/grpc"
)

type HttpRepository interface {
	CreateUser(context.Context, entities.User) (int32, error)
}

type httpRespository struct {
	client pb.UserServiceClient
	log    log.Logger
}

func NewHttpRespository(conn *grpc.ClientConn, log log.Logger) *httpRespository {
	return &httpRespository{
		client: pb.NewUserServiceClient(conn),
		log:    log,
	}
}

func (r *httpRespository) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	usrReq := &pb.CreateUserRequest{User: &pb.User{PwdHash: usr.Pwd_hash, Name: usr.Name, Age: usr.Age, AdditionalInformation: usr.Additional_information, Parent: usr.Parent}}
	userResponse, err := r.client.CreateUser(ctx, usrReq)
	if err != nil {
		level.Error(r.log).Log("Client error creating user", err)
		return 0, errors.NewServiceResponseError()
	}
	return userResponse.User.Id, nil
}
