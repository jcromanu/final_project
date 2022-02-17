package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
)

type CustomError interface {
	Code() int
	Error() string
}

type HttpError struct {
	StatusCode int
	Message    string
}

func (err *HttpError) Error() string {
	return err.Message
}

func (err *HttpError) Code() int {
	return err.StatusCode
}

func UnexpectedSigningMethod() error {
	return errors.New("unexpected signing method")
}

func NewBadRequestError() error {
	return &HttpError{StatusCode: 400, Message: "Bad request"}
}

func NewEmptyFieldError() error {
	return &HttpError{StatusCode: 400, Message: "Bad request , empty field "}
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
