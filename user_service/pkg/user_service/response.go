package userservice

import "github.com/jcromanu/final_project/user_service/pkg/entities"

type createUserResponse struct {
	User    entities.User
	Message entities.Message
}
