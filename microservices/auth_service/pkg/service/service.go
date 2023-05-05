package service

import (
	"context"
	"errors"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/auth_service/pkg/db"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
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
	db     db.UserStore
	logger log.Logger
}

func NewAuthSvc(db db.UserStore, logger log.Logger) AuthSvc {
	return AuthSvc{db: db, logger: logger}
}

func (s AuthSvc) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	level.Info(s.logger).Log("msg", "Attempting login for user", "username", req.Username)
	user, err := s.db.GetUser(req.Username)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error getting user", "error", err)
		return nil, ErrInvalidCredentials
	}
	level.Info(s.logger).Log("msg", "User retrieved successfully", "username", user.Username)
	level.Info(s.logger).Log("msg", "Comparing passwords", "DB password:", user.Password, " request password:", req.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		level.Error(s.logger).Log("msg", "Invalid password", "error", err)
		return nil, ErrInvalidCredentials
	}

	token, err := s.generateToken(user.Id, user.Username)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error generating token", "error", err)
		return nil, ErrInternalServerError
	}

	level.Info(s.logger).Log("msg", "Login successful for user", "username", req.Username)
	return &auth.LoginResponse{
		Token:   token,
		Message: "login successful",
	}, nil
}

func (s AuthSvc) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	level.Info(s.logger).Log("msg", "Getting user by ID", "id", req.Id)
	user, err := s.db.GetUserByID(req.Id)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error getting user by ID", "error", err)
		return nil, ErrUnauthorized
	}

	level.Info(s.logger).Log("msg", "User retrieved successfully", "username", user.Username)
	return &auth.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s AuthSvc) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	level.Info(s.logger).Log("msg", "Creating user", "username", req.Username)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error generating hashed password", "error", err)
		return nil, ErrInternalServerError
	}

	user := auth.User{
		Username: req.Username,
		Password: string(hashedPwd),
		Email:    req.Email,
	}

	if err := s.db.CreateUser(&user); err != nil {
		level.Error(s.logger).Log("msg", "Error creating user", "error", err)
		return nil, ErrInternalServerError
	}

	level.Info(s.logger).Log("msg", "User created successfully", "username", req.Username)
	return &auth.CreateUserResponse{
		Id:      user.Id,
		Message: "user created successfully",
	}, nil
}

func (s AuthSvc) generateToken(userID, username string) (string, error) {
	level.Info(s.logger).Log("msg", "Generating token for user", "username", username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userID,
		"username": username,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		level.Error(s.logger).Log("msg", "Error signing token", "error", err)
		return "", err
	}

	level.Info(s.logger).Log("msg", "Token generated successfully for user", "username", username)
	return tokenString, nil
}
