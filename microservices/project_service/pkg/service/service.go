package service

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/db"
	"github.com/go-kit/kit/log"
)

// ProjectService is a struct that implements the project.UnimplementedProjectServiceServer interface.
// It also contains a database interface and a logger.
type ProjectService struct {
	project.UnimplementedProjectServiceServer
	DB     db.ProjectDB
	Logger log.Logger
}

// NewProjectService is a constructor function that takes a database interface and a logger and
// returns a new instance of ProjectService.
func NewProjectService(database db.ProjectDB, logger log.Logger) *ProjectService {
	return &ProjectService{
		DB:     database,
		Logger: logger,
	}
}

// GetProject is a method that retrieves a project by its ID.
// It logs the error if there is one, and it also logs the successful retrieval of a project.
func (s *ProjectService) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	p, err := s.DB.GetProjectByID(req.Id)
	if err != nil {
		s.Logger.Log("error", err)
		return nil, err
	}

	s.Logger.Log("info", "GetProject successful", "id", p.ID)

	return &project.GetProjectResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}, nil
}

// CreateProject is a method that creates a new project.
// It logs the error if there is one, and it also logs the successful creation of a project.
func (s *ProjectService) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	dbproject := &db.Project{
		ID:          "some_generated_id",
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.DB.CreateProject(dbproject)
	if err != nil {
		s.Logger.Log("error", err)
		return nil, err
	}

	s.Logger.Log("info", "CreateProject successful", "id", dbproject.ID)

	return &project.CreateProjectResponse{
		Id:      dbproject.ID,
		Message: "Project created successfully",
	}, nil
}

// DeleteProject is a method that deletes a project by its ID.
// It logs the error if there is one, and it also logs the successful deletion of a project.
func (s *ProjectService) DeleteProject(ctx context.Context, req *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	err := s.DB.DeleteProject(req.Id)
	if err != nil {
		s.Logger.Log("error", err)
		return nil, err
	}

	s.Logger.Log("info", "DeleteProject successful", "id", req.Id)

	return &project.DeleteProjectResponse{
		Message: "Project deleted successfully",
	}, nil
}

// GetProjects retrieves all projects.
// It logs the error if there is one, and it also logs the successful retrieval of the projects.
func (s *ProjectService) GetProjects(ctx context.Context, req *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {
	projects, err := s.DB.GetProjects()
	if err != nil {
		s.Logger.Log("error", err)
		return nil, err
	}

	s.Logger.Log("info", "GetProjects successful")

	responseProjects := make([]*project.Project, len(projects))
	for i, p := range projects {
		responseProjects[i] = &project.Project{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
		}
	}

	return &project.GetProjectsResponse{
		Projects: responseProjects,
	}, nil
}
