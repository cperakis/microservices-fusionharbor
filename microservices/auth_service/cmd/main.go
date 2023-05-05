package main

import (
	"log"
	"net"

	"github.com/fusionharbor/microservices/auth_service/pkg/db"
	"github.com/fusionharbor/microservices/auth_service/pkg/endpoints"
	"github.com/fusionharbor/microservices/auth_service/pkg/service"
	"github.com/fusionharbor/microservices/auth_service/pkg/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	dsn := "root:@/fusionharbor?charset=utf8&parseTime=True&loc=Local"
	userStore, err := db.NewGormUserStore(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	authService := service.NewAuthSvc(userStore)
	endpoints := endpoints.NewEndpoints(authService)
	transport.RegisterAuthGRPCServer(grpc.ServiceRegistrar(s), endpoints)

	log.Println("Starting auth-service on port 8081...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
