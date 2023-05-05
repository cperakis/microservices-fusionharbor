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

type AuthProxy struct {
	client auth.AuthServiceClient
	logger log.Logger
}

func NewAuthProxy(client auth.AuthServiceClient, logger log.Logger) *AuthProxy {
	return &AuthProxy{
		client: client,
		logger: logger,
	}
}

func (p *AuthProxy) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", p.login).Methods("POST")
	r.HandleFunc("/api/user/{id}", p.getUser).Methods("GET")
	r.HandleFunc("/api/user", p.createUser).Methods("POST")
}

func (p *AuthProxy) login(w http.ResponseWriter, r *http.Request) {
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

func (p *AuthProxy) getUser(w http.ResponseWriter, r *http.Request) {
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

func (p *AuthProxy) createUser(w http.ResponseWriter, r *http.Request) {
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
