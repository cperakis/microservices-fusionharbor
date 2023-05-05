package transport

import (
	"context"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/endpoints"
	"github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getProject    grpc.Handler
	createProject grpc.Handler
	project.UnimplementedProjectServiceServer
}

func (s *grpcServer) mustEmbedUnimplementedProjectServiceServer() {}

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
	}
}

func (s *grpcServer) GetProject(ctx context.Context, req *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	_, resp, err := s.getProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.GetProjectResponse), nil
}

func (s *grpcServer) CreateProject(ctx context.Context, req *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	_, resp, err := s.createProject.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*project.CreateProjectResponse), nil
}

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
