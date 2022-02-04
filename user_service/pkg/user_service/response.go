package userservice

import "github.com/jcromanu/final_project/user_service/pkg/entities"

type createUserResponse struct {
	User    entities.User
	Message entities.Message
}

type getUserResponse struct {
	User    entities.User
	Message entities.Message
}

type updateUserResponse struct {
	Message entities.Message
}

type deleteUserResponse struct {
	Message entities.Message
}
