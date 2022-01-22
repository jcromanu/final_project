package errors

import "errors"

func NewInternalError() error {
	return errors.New("Internal service error")
}

func NewDatabaseError() error {
	return errors.New("Database error")
}
