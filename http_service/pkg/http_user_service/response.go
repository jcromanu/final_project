package httpuserservice

import "github.com/jcromanu/final_project/http_service/pkg/entities"

type createUserResponse struct {
	User    entities.User
	Message entities.Message
}
