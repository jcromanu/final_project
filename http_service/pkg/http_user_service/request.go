package httpuserservice

import "github.com/jcromanu/final_project/http_service/pkg/entities"

type createUserRequest struct {
	User entities.User
}

type getUserRequest struct {
	Id int32
}

type updateUserRequest struct {
	User entities.User
}

type deleteUserRequest struct {
	Id int32
}
