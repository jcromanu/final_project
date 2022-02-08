package httpuserservice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func decodePostCreateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, err
	}
	request = getUserRequest{Id: int32(id)}
	return request, nil
}

func decodePutUpdateUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req updateUserRequest
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	req.User.Id = int32(id)
	return req, nil
}

func decodeDeleteUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, err
	}
	request = deleteUserRequest{Id: int32(id)}
	return request, nil
}

type errorer interface {
	error() error
}
