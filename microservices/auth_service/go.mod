module github.com/fusionharbor/microservices/auth_service

replace github.com/fusionharbor/microservices/api => ../api

replace github.com/fusionharbor/microservices/auth_service => ../auth_service

go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fusionharbor/microservices/api v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.12.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.0
	google.golang.org/grpc v1.54.0
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.25.0
)

require golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa

require (
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
