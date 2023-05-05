// Package main contains the entry point for the auth_service microservice.
package main

import (
	"net"
	"os"

	// Importing packages for database, endpoints, services, and transport
	"github.com/fusionharbor/microservices/auth_service/pkg/db"
	"github.com/fusionharbor/microservices/auth_service/pkg/endpoints"
	"github.com/fusionharbor/microservices/auth_service/pkg/service"
	"github.com/fusionharbor/microservices/auth_service/pkg/transport"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create a logger instance
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())

	// Listen for incoming connections on port 8081
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		level.Error(logger).Log("msg", "Failed to listen", "error", err)
		os.Exit(1)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the reflection service on the server to support debugging tools
	reflection.Register(s)

	// Configure the database connection
	dsn := "cperakis:@/fusionharbor?charset=utf8&parseTime=True&loc=Local"

	// Create a new GormUserStore with the given data source
	userStore, err := db.NewGormUserStore(dsn)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to connect to the database", "error", err)
		os.Exit(1)
	}

	// Create a new auth service instance with the user store and logger
	authService := service.NewAuthSvc(userStore, logger)

	// Create the endpoints for the auth service
	endpoints := endpoints.NewEndpoints(authService)

	// Register the AuthGRPCServer with the gRPC server
	transport.RegisterAuthGRPCServer(grpc.ServiceRegistrar(s), endpoints)

	// Start the server and listen for incoming requests
	level.Info(logger).Log("msg", "Starting auth-service on port 8081...")
	if err := s.Serve(listener); err != nil {
		level.Error(logger).Log("msg", "Failed to serve", "error", err)
		os.Exit(1)
	}
}
