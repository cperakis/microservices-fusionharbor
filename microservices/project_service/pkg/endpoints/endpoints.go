package endpoints

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetProject    endpoint.Endpoint
	CreateProject endpoint.Endpoint
}

func MakeEndpoints(s service.ProjectService) Endpoints {
	return Endpoints{
		GetProject:    makeGetProjectEndpoint(s),
		CreateProject: makeCreateProjectEndpoint(s),
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
