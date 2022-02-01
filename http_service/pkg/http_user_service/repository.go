package httpuserservice

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jcromanu/final_project/http_service/pkg/entities"
	"github.com/jcromanu/final_project/user_service/pb"
	"google.golang.org/grpc"
)

type repository struct {
	client pb.UserServiceClient
	log    log.Logger
}

func NewRespository(conn *grpc.ClientConn, log log.Logger) *repository {
	return &repository{
		client: pb.NewUserServiceClient(conn),
		log:    log,
	}
}

func (r *repository) CreateUser(ctx context.Context, usr entities.User) (int32, error) {
	usrReq := &pb.CreateUserRequest{User: &pb.User{PwdHash: usr.Pwd_hash, Name: usr.Name, Age: usr.Age, AdditionalInformation: usr.Additional_information, Parent: usr.Parent}}
	userResponse, err := r.client.CreateUser(ctx, usrReq)
	if err != nil {
		level.Error(r.log).Log("Client error creating user", err)
		return 0, err
	}
	return userResponse.User.Id, nil
}

func (r *repository) GetUser(ctx context.Context, id int32) (entities.User, error) {
	usrReq := &pb.GetUserRequest{Id: id}
	userResponse, err := r.client.GetUser(ctx, usrReq)
	if err != nil {
		level.Error(r.log).Log("Client error getting user", err)
		return entities.User{}, err
	}
	usr := entities.User{
		Id:                     userResponse.User.Id,
		Age:                    userResponse.User.Age,
		Additional_information: userResponse.User.AdditionalInformation,
		Pwd_hash:               userResponse.User.AdditionalInformation,
		Name:                   userResponse.User.Name,
	}
	return usr, nil
}
