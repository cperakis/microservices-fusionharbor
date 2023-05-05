package service

import (
	"context"
	"errors"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/auth_service/pkg/db"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)

type AuthSvc struct {
	auth.UnimplementedAuthServiceServer
	db db.UserStore
}

func NewAuthSvc(db db.UserStore) AuthSvc {
	return AuthSvc{db: db}
}

func (s AuthSvc) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := s.db.GetUser(req.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := generateToken(user.Id, user.Username)
	if err != nil {
		return nil, ErrInternalServerError
	}

	return &auth.LoginResponse{
		Token:   token,
		Message: "login successful",
	}, nil
}

func (s AuthSvc) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	user, err := s.db.GetUserByID(req.Id)
	if err != nil {
		return nil, ErrUnauthorized
	}

	return &auth.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s AuthSvc) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrInternalServerError
	}

	user := auth.User{
		Username: req.Username,
		Password: string(hashedPwd),
		Email:    req.Email,
	}

	if err := s.db.CreateUser(&user); err != nil {
		return nil, ErrInternalServerError
	}

	return &auth.CreateUserResponse{
		Id:      user.Id,
		Message: "user created successfully",
	}, nil
}

func generateToken(userID, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userID,
		"username": username,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
