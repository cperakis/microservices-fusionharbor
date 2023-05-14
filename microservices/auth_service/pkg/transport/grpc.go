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
	getUsers   grpc.Handler
	createUser grpc.Handler
	deleteUser grpc.Handler
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
		deleteUser: grpc.NewServer(
			endpoints.DeleteUserEndpoint,
			DecodeGRPCDeleteUserRequest,
			EncodeGRPCDeleteUserResponse,
		),
		getUsers: grpc.NewServer(
			endpoints.GetUsersEndpoint,
			DecodeGRPCGetUsersRequest,
			EncodeGRPCGetUsersResponse,
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

func (s *grpcServer) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (*auth.DeleteUserResponse, error) {
	_, resp, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*auth.DeleteUserResponse), nil
}

func (s *grpcServer) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	_, resp, err := s.getUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*auth.GetUsersResponse), nil
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

func DecodeGRPCGetUsersRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCGetUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeGRPCDeleteUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func EncodeGRPCDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
