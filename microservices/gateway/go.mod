module github.com/cperakis/fusionharbor/pb

replace github.com/fusionharbor/microservices/api => ../api

replace github.com/fusionharbor/microservices/auth_service => ../auth_service

replace github.com/fusionharbor/microservices/project_service => ../project_service

replace github.com/fusionharbor/microservices/gateway => ../gateway

go 1.17

require (
	github.com/fusionharbor/microservices/api v0.0.0-00010101000000-000000000000
	github.com/fusionharbor/microservices/gateway v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0
	github.com/gorilla/mux v1.8.0
	github.com/pelletier/go-toml v1.9.5
	google.golang.org/grpc v1.54.0
)

require (
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
