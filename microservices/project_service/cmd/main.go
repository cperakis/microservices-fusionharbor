package main

import (
	"fmt"
	"net"
	"os"

	"github.com/fusionharbor/microservices/project_service/confs"

	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/project_service/pkg/db"
	"github.com/fusionharbor/microservices/project_service/pkg/endpoints"
	"github.com/fusionharbor/microservices/project_service/pkg/service"
	"github.com/fusionharbor/microservices/project_service/pkg/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
	// Create a logger instance
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())

	// Set up different log levels
	infoLogger := level.Info(logger)
	errorLogger := level.Error(logger)

	infoLogger.Log("message", "Starting Project Service") // Log an informational message

	listener, err := net.Listen("tcp", confs.Conf.Port)
	if err != nil {
		errorLogger.Log("error", err) // Log any errors with the error level
		panic(err)
	}

	dsn := confs.Conf.Database
	projectDB, err := db.NewGormProjectDB(dsn)
	if err != nil {
		errorLogger.Log("error", fmt.Sprintf("Failed to connect to the database: %v", err))
		panic(err)
	}

	projectService := service.NewProjectService(projectDB, logger)

	projectEndpoints := endpoints.MakeEndpoints(*projectService)

	grpcServer := grpc.NewServer()
	project.RegisterProjectServiceServer(grpcServer, transport.NewGRPCServer(projectEndpoints))

	infoLogger.Log("message", "Starting Project Service on :8082") // Log an informational message
	grpcServer.Serve(listener)
}
