package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(e endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/project/{id}").Handler(httptransport.NewServer(
		e.GetProject,
		decodeGetProjectRequest,
		encodeGetProjectResponse,
	))

	r.Methods("POST").Path("/project").Handler(httptransport.NewServer(
		e.CreateProject,
		decodeCreateProjectRequest,
		encodeCreateProjectResponse,
	))

	return r
}

func decodeGetProjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return &project.GetProjectRequest{Id: id}, nil
}

func encodeGetProjectResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*project.GetProjectResponse)
	return json.NewEncoder(w).Encode(res)
}

func decodeCreateProjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req project.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeCreateProjectResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*project.CreateProjectResponse)
	return json.NewEncoder(w).Encode(res)
}

var ErrBadRouting = errors.New("invalid route")
