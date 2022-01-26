package userservice

import "github.com/jcromanu/final_project/user_service/pkg/entities"

/*
type authenticateResponse struct {
	MessageResponse
	token string
}*/

type createUserResponse struct {
	User    entities.User
	Message entities.Message
}

/*
type getUserResponse struct {
	MessageResponse
	entities.User
}

type updateUserResponse struct {
	MessageResponse
}

type deleteUserResponse struct {
	MessageResponse
}
*/
