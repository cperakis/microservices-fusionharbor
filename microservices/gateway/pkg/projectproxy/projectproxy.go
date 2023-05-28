package projectproxy

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/gorilla/mux"
)

type ProjectProxy struct {
	client project.ProjectServiceClient
}

func NewProjectProxy(client project.ProjectServiceClient) *ProjectProxy {
	return &ProjectProxy{
		client: client,
	}
}

func (p *ProjectProxy) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/project/{id}", p.getProject).Methods("GET")
	r.HandleFunc("/api/projects", p.getProjects).Methods("GET")
	r.HandleFunc("/api/project", p.createProject).Methods("POST")
	r.HandleFunc("/api/project/{id}", p.deleteProject).Methods("DELETE")
}

func (p *ProjectProxy) getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "invalid route", http.StatusBadRequest)
		return
	}

	req := &project.GetProjectRequest{Id: id}
	res, err := p.client.GetProject(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (p *ProjectProxy) deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	req := &project.DeleteProjectRequest{Id: id}
	res, err := p.client.DeleteProject(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (p *ProjectProxy) getProjects(w http.ResponseWriter, r *http.Request) {

	req := &project.GetProjectsRequest{}
	res, err := p.client.GetProjects(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (p *ProjectProxy) createProject(w http.ResponseWriter, r *http.Request) {
	var req project.CreateProjectRequest
	log.Println("createProject request received", req.Name, req.Description, req.Metadata)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := p.client.CreateProject(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
