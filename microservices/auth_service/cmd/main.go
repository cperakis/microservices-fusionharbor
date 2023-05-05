// Package main contains the entry point for the auth_service microservice.
package main

import (
	"log"
	"net"

	// Importing packages for database, endpoints, services, and transport
	"github.com/fusionharbor/microservices/auth_service/pkg/db"
	"github.com/fusionharbor/microservices/auth_service/pkg/endpoints"
	"github.com/fusionharbor/microservices/auth_service/pkg/service"
	"github.com/fusionharbor/microservices/auth_service/pkg/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Listen for incoming connections on port 8081
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
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
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create a new auth service instance with the user store
	authService := service.NewAuthSvc(userStore)

	// Create the endpoints for the auth service
	endpoints := endpoints.NewEndpoints(authService)

	// Register the AuthGRPCServer with the gRPC server
	transport.RegisterAuthGRPCServer(grpc.ServiceRegistrar(s), endpoints)

	// Start the server and listen for incoming requests
	log.Println("Starting auth-service on port 8081...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
