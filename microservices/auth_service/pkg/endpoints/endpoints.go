package endpoints

import (
	"context"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/auth_service/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	LoginEndpoint      endpoint.Endpoint
	GetUserEndpoint    endpoint.Endpoint
	CreateUserEndpoint endpoint.Endpoint
}

func MakeLoginEndpoint(svc service.AuthSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*auth.LoginRequest)
		loginReq := &auth.LoginRequest{
			Username: req.Username,
			Password: req.Password,
		}
		token, err := svc.Login(ctx, loginReq)
		if err != nil {
			return nil, err
		}
		return &auth.LoginResponse{Token: token.Token}, nil
	}
}

func MakeGetUserEndpoint(svc service.AuthSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*auth.GetUserRequest)
		user, err := svc.GetUser(ctx, &auth.GetUserRequest{Id: req.Id, Token: req.Token})
		if err != nil {
			return nil, err
		}
		return &auth.GetUserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		}, nil
	}
}

func MakeCreateUserEndpoint(svc service.AuthSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*auth.CreateUserRequest)
		createUser := &auth.CreateUserRequest{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}
		user, err := svc.CreateUser(ctx, createUser)
		if err != nil {
			return nil, err
		}
		return &auth.CreateUserResponse{Id: user.Id}, nil
	}
}

func NewEndpoints(svc service.AuthSvc) Endpoints {
	return Endpoints{
		LoginEndpoint:      MakeLoginEndpoint(svc),
		GetUserEndpoint:    MakeGetUserEndpoint(svc),
		CreateUserEndpoint: MakeCreateUserEndpoint(svc),
	}
}
