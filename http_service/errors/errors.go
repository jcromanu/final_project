package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
)

func NewInternalError() error {
	return errors.New("Internal service error")
}

func NewDatabaseError() error {
	return errors.New("Database error")
}

func NewBadTypeError() error {
	return errors.New("Bad type error ")
}

func NewBadResponseTypeError() error {
	return errors.New("Bad response type error ")
}

func NewBadRequestError() error {
	return errors.New("Bad request for method ")
}

func NewParsingRequestError() error {
	return errors.New("Error parsing user request ")
}

func NewServiceResponseError() error {
	return errors.New("Failed service response  ")
}

func NewErrBadRouting() error {
	return errors.New("inconsistent mapping between route and handler (programmer error)")
}

func ResolveHTTP(c codes.Code) int {
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

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)
