package httpuserservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jcromanu/final_project/http_service/errors"
)

func NewHTTPServer(e Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	opt := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		e.CreateUser,
		decodePostCreateUserRequest,
		encodeResponse,
		opt...,
	))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrAlreadyExists, errors.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
