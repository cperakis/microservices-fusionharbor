package service

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/db"
)

type ProjectService struct {
	project.UnimplementedProjectServiceServer
	DB db.ProjectDB
}

func NewProjectService(database db.ProjectDB) *ProjectService {
	return &ProjectService{
		DB: database,
	}
}

func (s *ProjectService) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	p, err := s.DB.GetProjectByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &project.GetProjectResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}, nil
}

func (s *ProjectService) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	dbproject := &db.Project{
		ID:          "some_generated_id",
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.DB.CreateProject(dbproject)
	if err != nil {
		return nil, err
	}

	return &project.CreateProjectResponse{
		Id:      dbproject.ID,
		Message: "Project created successfully",
	}, nil
}
