package userservice

import "github.com/jcromanu/final_project/pkg/entities"

type createUserRequest struct {
	User entities.User
}

/*
type updateUserRequest struct {
	entities.User
}

type getUserRequest struct {
	id int32
}

type deleteUserRequest struct {
	id int32
}

type authenticateRequest struct {
	token string
}
*/
