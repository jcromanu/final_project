package userservice

import "github.com/jcromanu/final_project/user_service/pkg/entities"

type createUserRequest struct {
	user entities.User
}

type getUserRequest struct {
	id int32
}

type updateUserRequest struct {
	user entities.User
}

type deleteUserRequest struct {
	id int32
}
