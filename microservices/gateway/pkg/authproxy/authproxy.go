package authproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/gorilla/mux"
)

type AuthProxy struct {
	client auth.AuthServiceClient
}

func NewAuthProxy(client auth.AuthServiceClient) *AuthProxy {
	return &AuthProxy{
		client: client,
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
		return
	}

	req := &auth.GetUserRequest{Id: id}
	res, err := p.client.GetUser(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
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
