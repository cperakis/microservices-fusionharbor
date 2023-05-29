package confs

import (
	"log"

	toml "github.com/pelletier/go-toml"
)

type Configuration struct {
	ExpirationCookie int // In seconds
	Database         string
	Port             string
	AuthPort         string
	ProjectPort      string
}

var Conf Configuration

func init() {
	config, err := toml.LoadFile("../confs/conf_gateway.toml")
	if err != nil {
		log.Fatal(err)
	}

	expirationCookie, ok := config.Get("expiration_cookie").(int64)
	if !ok {
		Conf.ExpirationCookie = 7200
	} else {
		Conf.ExpirationCookie = int(expirationCookie)
	}

	database, ok := config.Get("database").(string)
	if ok {
		Conf.Database = database
	} else {
		Conf.Database = "cperakis:@/fusionharbor?charset=utf8&parseTime=True&loc=Local"
	}

	port, ok := config.Get("port").(string)
	if ok {
		Conf.Port = port
	} else {
		Conf.Port = ":8080"
	}

	authPort, ok := config.Get("auth_port").(string)
	if ok {
		Conf.AuthPort = authPort
	} else {
		Conf.AuthPort = ":8081"
	}

	projectPort, ok := config.Get("project_port").(string)
	if ok {
		Conf.ProjectPort = projectPort
	} else {
		Conf.ProjectPort = ":8082"
	}

}
