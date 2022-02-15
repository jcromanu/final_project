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
	usrReq := &pb.CreateUserRequest{
		User: &pb.User{
			PwdHash:               usr.PwdHash,
			Name:                  usr.Name,
			Age:                   usr.Age,
			AdditionalInformation: usr.AdditionalInformation,
			Parent:                usr.Parent,
			Email:                 usr.Email}}
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
		Id:                    userResponse.User.Id,
		Age:                   userResponse.User.Age,
		AdditionalInformation: userResponse.User.AdditionalInformation,
		PwdHash:               userResponse.User.PwdHash,
		Name:                  userResponse.User.Name,
		Parent:                userResponse.User.Parent,
		Email:                 userResponse.User.Email,
	}
	return usr, nil
}

func (r *repository) UpdateUser(ctx context.Context, usr entities.User) (string, error) {
	usrReq := &pb.UpdateUserRequest{
		User: &pb.User{
			Id:                    usr.Id,
			PwdHash:               usr.PwdHash,
			Name:                  usr.Name,
			Age:                   usr.Age,
			AdditionalInformation: usr.AdditionalInformation,
			Parent:                usr.Parent,
			Email:                 usr.Email}}
	resp, err := r.client.UpdateUser(ctx, usrReq)
	if err != nil {
		level.Error(r.log).Log("Client error updating user ", err)
		return "", err
	}
	return resp.Message.Message, nil
}

func (r *repository) DeleteUser(ctx context.Context, id int32) (string, error) {
	userReq := &pb.DeleteUserRequest{Id: id}
	resp, err := r.client.DeleteUser(ctx, userReq)
	if err != nil {
		level.Error(r.log).Log("Client error deleting  user ", err)
		return "", err
	}
	return resp.Message.Message, nil
}
