package endpoints

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the project service.
type Endpoints struct {
	GetProject    endpoint.Endpoint
	CreateProject endpoint.Endpoint
	DeleteProject endpoint.Endpoint
	GetProjects   endpoint.Endpoint // New endpoint for getting all projects
}

// MakeEndpoints initializes all Go kit endpoints for the project service.
func MakeEndpoints(s service.ProjectService) Endpoints {
	return Endpoints{
		GetProject:    makeGetProjectEndpoint(s),
		CreateProject: makeCreateProjectEndpoint(s),
		DeleteProject: makeDeleteProjectEndpoint(s),
		GetProjects:   makeGetProjectsEndpoint(s), // Initialize new endpoint for getting all projects
	}
}

func makeGetProjectEndpoint(s service.ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*project.GetProjectRequest)
		res, err := s.GetProject(ctx, req)
		return res, err
	}
}

func makeCreateProjectEndpoint(s service.ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*project.CreateProjectRequest)
		res, err := s.CreateProject(ctx, req)
		return res, err
	}
}

func makeDeleteProjectEndpoint(s service.ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*project.DeleteProjectRequest)
		res, err := s.DeleteProject(ctx, req)
		return res, err
	}
}

func makeGetProjectsEndpoint(s service.ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*project.GetProjectsRequest)
		res, err := s.GetProjects(ctx, req)
		return res, err
	}
}
