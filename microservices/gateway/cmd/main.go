package main

import (
	"net/http"
	"os"

	"github.com/fusionharbor/microservices/gateway/confs"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/gateway/pkg/authproxy"
	"github.com/fusionharbor/microservices/gateway/pkg/projectproxy"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	// Create gRPC connections to Auth and Project services
	authConn, err := grpc.Dial(confs.Conf.AuthPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer authConn.Close()

	projectConn, err := grpc.Dial(confs.Conf.ProjectPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer projectConn.Close()

	// Create gRPC clients
	authClient := auth.NewAuthServiceClient(authConn)
	projectClient := project.NewProjectServiceClient(projectConn)

	// Create Auth and Project proxies
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowInfo())
	authProxy := authproxy.NewAuthProxy(authClient, logger)
	projectProxy := projectproxy.NewProjectProxy(projectClient)

	// Create the gateway mux
	r := mux.NewRouter()
	authProxy.RegisterRoutes(r)
	projectProxy.RegisterRoutes(r)

	// Start the gateway
	level.Info(logger).Log("msg", "Starting gateway service on port 8080...")
	if err := http.ListenAndServe(confs.Conf.Port, r); err != nil {
		panic(err)
	}
}
