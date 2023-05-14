package transport

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/endpoints"
	"github.com/go-kit/kit/transport/grpc"
)

// grpcServer wraps endpoints in a gRPC server.
type grpcServer struct {
	getProject    grpc.Handler
	createProject grpc.Handler
	deleteProject grpc.Handler
	getProjects   grpc.Handler // New handler for getting all projects
	project.UnimplementedProjectServiceServer
}

// NewGRPCServer creates a new gRPC server that wraps the given endpoints.
func NewGRPCServer(endpoints endpoints.Endpoints) project.ProjectServiceServer {
	return &grpcServer{
		getProject: grpc.NewServer(
			endpoints.GetProject,
			DecodeGRPCGetProjectRequest,
			EncodeGRPCGetProjectResponse,
		),
		createProject: grpc.NewServer(
			endpoints.CreateProject,
			DecodeGRPCCreateProjectRequest,
			EncodeGRPCCreateProjectResponse,
		),
		deleteProject: grpc.NewServer(
			endpoints.DeleteProject,
			DecodeGRPCDeleteProjectRequest,
			EncodeGRPCDeleteProjectResponse,
		),
		getProjects: grpc.NewServer(
			endpoints.GetProjects,         // New endpoint for getting all projects
			DecodeGRPCGetProjectsRequest,  // New decode request function
			EncodeGRPCGetProjectsResponse, // New encode response function
		),
	}
}

// GetProject is a method that serves a gRPC request to get a project.
func (s *grpcServer) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	_, resp, err := s.getProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.GetProjectResponse), nil
}

// CreateProject is a method that serves a gRPC request to create a project.
func (s *grpcServer) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	_, resp, err := s.createProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.CreateProjectResponse), nil
}

// DeleteProject is a method that serves a gRPC request to delete a project.
func (s *grpcServer) DeleteProject(ctx context.Context, req *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	_, resp, err := s.deleteProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.DeleteProjectResponse), nil
}

// GetProjects is a method that serves a gRPC request to get all projects. // New function
func (s *grpcServer) GetProjects(ctx context.Context, req *project.GetProjectsRequest) (*project.GetProjectsResponse, error) {
	_, resp, err := s.getProjects.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.GetProjectsResponse), nil
}

// Decode and encode functions for each endpoint

func DecodeGRPCGetProjectRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCGetProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCCreateProjectRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCCreateProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCDeleteProjectRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCDeleteProjectResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGRPCGetProjectsRequest is a function that decodes a gRPC request into a GetProjectsRequest.
func DecodeGRPCGetProjectsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

// EncodeGRPCGetProjectsResponse is a function that encodes a GetProjectsResponse into a gRPC response.
func EncodeGRPCGetProjectsResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
