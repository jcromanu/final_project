package errors

import "errors"

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

func Error() error {
	return errors.New("Error  ")
}
