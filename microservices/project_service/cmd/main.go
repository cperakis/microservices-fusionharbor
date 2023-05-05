package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/db"
	"github.com/fusionharbor/microservices/project_service/pkg/endpoints"
	"github.com/fusionharbor/microservices/project_service/pkg/service"
	"github.com/fusionharbor/microservices/project_service/pkg/transport"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	dsn := "youruser:yourpassword@tcp(localhost:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	projectDB, err := db.NewGormProjectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	projectService := service.NewProjectService(projectDB)

	projectEndpoints := endpoints.MakeEndpoints(*projectService)

	grpcServer := grpc.NewServer()
	project.RegisterProjectServiceServer(grpcServer, transport.NewGRPCServer(projectEndpoints))

	fmt.Println("Starting Project Service on :8082")
	grpcServer.Serve(listener)
}
