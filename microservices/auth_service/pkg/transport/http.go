package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/auth_service/pkg/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(e endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		e.LoginEndpoint,
		decodeLoginRequest,
		encodeLoginResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		e.GetUserEndpoint,
		decodeGetUserRequest,
		encodeGetUserResponse,
	))

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		e.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	))

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req auth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*auth.LoginResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return &auth.GetUserRequest{Id: id}, nil
}

func encodeGetUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*auth.GetUserResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req auth.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeCreateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*auth.CreateUserResponse)
	return json.NewEncoder(w).Encode(res)
}

var ErrBadRouting = errors.New("invalid route")
