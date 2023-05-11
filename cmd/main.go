package main

import (
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/service"
	"log"
)

var konfig *config.Config

func init() {
	var err error
	// load config
	konfig, err = config.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	// connect to Database
	err = db.ConnectDB(konfig)
	if err != nil {
		panic(err)
	}

	// connect to redis
	err = db.ConnectRedis(konfig)
	if err != nil {
		panic(err)
	}
}

func main() {

	log.Println("Starting Grpc Server")

	_, err := service.StartGrpcServer(konfig)

	if err != nil {
		log.Println(err)
	}

	log.Println("Grpc Server Running")

}
