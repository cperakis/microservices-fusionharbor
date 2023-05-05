package transport

import (
	"context"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/auth_service/pkg/endpoints"
	"github.com/go-kit/kit/transport/grpc"
	grpc2 "google.golang.org/grpc"
)

type grpcServer struct {
	login      grpc.Handler
	getUser    grpc.Handler
	createUser grpc.Handler
	auth.UnimplementedAuthServiceServer
}

func (s *grpcServer) mustEmbedUnimplementedAuthServiceServer() {}

func NewGRPCServer(endpoints endpoints.Endpoints) auth.AuthServiceServer {
	return &grpcServer{
		login: grpc.NewServer(
			endpoints.LoginEndpoint,
			DecodeGRPCLoginRequest,
			EncodeGRPCLoginResponse,
		),
		getUser: grpc.NewServer(
			endpoints.GetUserEndpoint,
			DecodeGRPCGetUserRequest,
			EncodeGRPCGetUserResponse,
		),
		createUser: grpc.NewServer(
			endpoints.CreateUserEndpoint,
			DecodeGRPCCreateUserRequest,
			EncodeGRPCCreateUserResponse,
		),
	}
}

func RegisterAuthGRPCServer(s grpc2.ServiceRegistrar, endpoints endpoints.Endpoints) {
	server := NewGRPCServer(endpoints)
	auth.RegisterAuthServiceServer(s, server)
}

func (s *grpcServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	_, resp, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*auth.LoginResponse), nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*auth.GetUserResponse), nil
}

func (s *grpcServer) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*auth.CreateUserResponse), nil
}

func DecodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCGetUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
