package main

import (
	"fmt"
	"net/http"

	"github.com/fusionharbor/microservices/api/auth"
	"github.com/fusionharbor/microservices/api/project"
	"github.com/fusionharbor/microservices/gateway/pkg/authproxy"
	"github.com/fusionharbor/microservices/gateway/pkg/projectproxy"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	// Create gRPC connections to Auth and Project services
	authConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer authConn.Close()

	projectConn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer projectConn.Close()

	// Create gRPC clients
	authClient := auth.NewAuthServiceClient(authConn)
	projectClient := project.NewProjectServiceClient(projectConn)

	// Create Auth and Project proxies
	authProxy := authproxy.NewAuthProxy(authClient)
	projectProxy := projectproxy.NewProjectProxy(projectClient)

	// Create the gateway mux
	r := mux.NewRouter()
	authProxy.RegisterRoutes(r)
	projectProxy.RegisterRoutes(r)

	// Start the gateway
	fmt.Println("Starting gateway on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
