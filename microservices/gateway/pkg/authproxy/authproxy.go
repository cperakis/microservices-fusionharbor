package authproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
)

// AuthProxy is a struct that holds the AuthServiceClient and a logger.
type AuthProxy struct {
	client auth.AuthServiceClient
	logger log.Logger
}

// NewAuthProxy initializes a new AuthProxy.
func NewAuthProxy(client auth.AuthServiceClient, logger log.Logger) *AuthProxy {
	return &AuthProxy{
		client: client,
		logger: logger,
	}
}

// RegisterRoutes registers the auth routes with the given router.
func (p *AuthProxy) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", p.login).Methods("POST")
	r.HandleFunc("/api/users/{id}", p.getUser).Methods("GET")
	r.HandleFunc("/api/users", p.getUsers).Methods("GET")
	r.HandleFunc("/api/users", p.createUser).Methods("POST")
}

// login handles the login endpoint.
func (p *AuthProxy) login(w http.ResponseWriter, r *http.Request) {
	level.Info(p.logger).Log("message", "Login request")
	var req auth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := p.client.Login(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// getUser handles the getUser endpoint.
func (p *AuthProxy) getUser(w http.ResponseWriter, r *http.Request) {
	level.Info(p.logger).Log("message", "GetUser request")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "invalid route", http.StatusBadRequest)
		level.Error(p.logger).Log("error", "invalid route")
		return
	}

	// Extract the token from the request header
	token := r.Header.Get("Authorization")
	level.Info(p.logger).Log("message", "GetUser request", "user_id", id, "token", token)
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		level.Error(p.logger).Log("error", "missing token")
		return
	}

	// Include the token in the GetUserRequest
	req := &auth.GetUserRequest{Id: id, Token: token}
	res, err := p.client.GetUser(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		level.Error(p.logger).Log("error", err.Error())
		return
	}

	json.NewEncoder(w).Encode(res)
	level.Info(p.logger).Log("message", "GetUser successful", "user_id", id)
}

// getUsers handles the getUsers endpoint.
func (p *AuthProxy) getUsers(w http.ResponseWriter, r *http.Request) {
	level.Info(p.logger).Log("message", "GetUsers request")
	// Extract the token from the request header
	token := r.Header.Get("Authorization")
	level.Info(p.logger).Log("message", "GetUsers request", "token", token)
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		level.Error(p.logger).Log("error", "missing token")
		return
	}

	// Include the token in the GetUserRequest
	req := &auth.GetUsersRequest{Token: token}
	res, err := p.client.GetUsers(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		level.Error(p.logger).Log("error", err.Error())
		return
	}

	json.NewEncoder(w).Encode(res)
	level.Info(p.logger).Log("message", "GetUsers successful")
}

// createUser handles the createUser endpoint.
func (p *AuthProxy) createUser(w http.ResponseWriter, r *http.Request) {
	level.Info(p.logger).Log("message", "CreateUser request")
	var req auth.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := p.client.CreateUser(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
