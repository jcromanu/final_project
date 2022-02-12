package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewUserNotFoundError() error {
	return status.Error(codes.NotFound, "User not found ")
}

func NewInternalError() error {
	return status.Error(codes.Internal, "Internal error")
}

func NewBadResponseTypeError() error {
	return status.Error(codes.Internal, "Malformed response  ")
}

func NewBadRequestError() error {
	return status.Error(codes.InvalidArgument, "Bad request")
}

func NewParsingRequestError() error {
	return errors.New("Error parsing user request ")
}

func NewProtoRequestError() error {
	return status.Error(codes.FailedPrecondition, "Proto request malformed")
}

func NewProtoResponseError() error {
	return status.Error(codes.Internal, "Proto response internal error ")
}

func Error() error {
	return errors.New("Error  ")
}
