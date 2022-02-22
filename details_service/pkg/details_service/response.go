package userservice

import "github.com/jcromanu/final_project/user_service/pkg/entities"

type createUserResponse struct {
	user    entities.User
	message entities.Message
}

type getUserResponse struct {
	user    entities.User
	message entities.Message
}

type updateUserResponse struct {
	message entities.Message
}

type deleteUserResponse struct {
	message entities.Message
}
