package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
)

func NewBadRequestError() error {
	return errors.New("Bad request for method ")
}

func NewEmptyFieldError() error {
	return errors.New("One or more fields must be not null")
}

func GrpcToHTTPCode(c codes.Code) int {
	switch c {
	case codes.FailedPrecondition:
		return 400
	case codes.Unauthenticated:
		return 401
	case codes.NotFound:
		return 404
	default:
		return 500
	}
}
