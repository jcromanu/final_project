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
	"google.golang.org/grpc/status"
)

func NewHTTPServer(e Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	opt := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Methods("POST").Path("/users").Handler(httptransport.NewServer(
		e.createUser,
		decodePostCreateUserRequest,
		encodeResponse,
		opt...,
	))
	r.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		e.getUser,
		decodeGetUserRequest,
		encodeResponse,
		opt...,
	))
	r.Methods("PUT").Path("/users/{id}").Handler(httptransport.NewServer(
		e.updateUser,
		decodePutUpdateUserRequest,
		encodeResponse,
		opt...,
	))
	r.Methods("DELETE").Path("/users/{id}").Handler(httptransport.NewServer(
		e.deleteUser,
		decodeDeleteUserRequest,
		encodeResponse,
		opt...,
	))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	e, ok := status.FromError(err)
	if !ok {
		e, ok := err.(errors.CustomError)
		if ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(e.Code())
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusText(http.StatusInternalServerError))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errors.GrpcToHTTPCode(e.Code()))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": e.Message(),
	})
}
