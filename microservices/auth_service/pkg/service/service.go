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

// GetUser retrieves user information by validating the token and fetching user details by ID.
func (s AuthSvc) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	level.Info(s.logger).Log("msg", "Getting user by ID", "id", req.Id)

	// Validate the JWT token.
	level.Info(s.logger).Log("msg", "Validating token: ", "token:", req.Token)
	token, err := s.validateToken(req.Token)
	if err != nil {
		level.Error(s.logger).Log("msg", "Invalid token", "error", err)
		return nil, ErrUnauthorized
	}

	// Print the raw token claims
	level.Info(s.logger).Log("msg", "Raw token claims:", "claims", token.Claims)

	// Get the user ID from the token claims.
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		level.Error(s.logger).Log("msg", "Failed to get claims from token")
		return nil, ErrUnauthorized
	}

	// Fetch user details by ID.
	user, err := s.db.GetUserByID(req.Id)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error getting user by ID", "error", err)
		return nil, ErrUnauthorized
	}

	level.Info(s.logger).Log("msg", "User retrieved successfully", "username", user.Username)
	return &auth.GetUserResponse{
		User: user,
	}, nil
}

// GetUser retrieves user information by validating the token and fetching user details by ID.
func (s AuthSvc) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) error {
	level.Info(s.logger).Log("msg", "Getting user by ID", "id", req.Id)

	// Validate the JWT token.
	level.Info(s.logger).Log("msg", "Validating token: ", "token:", req.Token)
	token, err := s.validateToken(req.Token)
	if err != nil {
		level.Error(s.logger).Log("msg", "Invalid token", "error", err)
		return ErrUnauthorized
	}

	// Print the raw token claims
	level.Info(s.logger).Log("msg", "Raw token claims:", "claims", token.Claims)

	// Get the user ID from the token claims.
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		level.Error(s.logger).Log("msg", "Failed to get claims from token")
		return ErrUnauthorized
	}

	// Fetch user details by ID.
	user, err := s.db.GetUserByID(req.Id)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error getting user by ID", "error", err)
		return ErrUnauthorized
	}

	s.db.DeleteUser(user.Id)
	if err != nil {
		level.Error(s.logger).Log("msg", "Error deleting user with ID", "error", err)
		return ErrUnauthorized
	}
	level.Info(s.logger).Log("msg", "User retrieved successfully", "username", user.Username)
	return nil
}

// GetUser retrieves user information by validating the token and fetching user details by ID.
func (s AuthSvc) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	level.Info(s.logger).Log("msg", "Getting users")
	// Validate the JWT token.
	level.Info(s.logger).Log("msg", "Validating token: ", "token:", req.Token)
	token, err := s.validateToken(req.Token)
	if err != nil {
		level.Error(s.logger).Log("msg", "Invalid token", "error", err)
		return nil, ErrUnauthorized
	}

	// Print the raw token claims
	level.Info(s.logger).Log("msg", "Raw token claims:", "claims", token.Claims)

	// Get the user ID from the token claims.
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		level.Error(s.logger).Log("msg", "Failed to get claims from token")
		return nil, ErrUnauthorized
	}

	// Fetch user details by ID.
	users, err := s.db.GetUsers()
	if err != nil {
		level.Error(s.logger).Log("msg", "Error getting users", "error", err)
		return nil, ErrUnauthorized
	}

	level.Info(s.logger).Log("msg", "Users retrieved successfully", users)
	return &auth.GetUsersResponse{
		Users: users,
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
		Role:     req.Role,
		Team:     req.Team,
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

// validateToken validates the JWT token and returns the parsed token if valid.
func (s AuthSvc) validateToken(tokenString string) (*jwt.Token, error) {
	level.Info(s.logger).Log("msg", "Validating token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnauthorized
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		level.Error(s.logger).Log("msg", "Invalid token", "error", err)
		return nil, ErrUnauthorized
	}

	level.Info(s.logger).Log("msg", "Token validated successfully:", token)
	return token, nil
}
