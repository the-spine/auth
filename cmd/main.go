package main

import (
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/events"
	"auth/internal/service"
	"log"
	"sync"
)

var (
	konfig *config.Config
	wg     sync.WaitGroup
)

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
		log.Println(err)
		panic(err)
	}

	// connect to redis
	err = db.ConnectRedis(konfig)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func main() {

	wg.Add(1)

	go func() {
		log.Println("Starting Grpc Server")

		defer wg.Done()

		_, err := service.StartGrpcServer(konfig)

		if err != nil {
			log.Println(err)
		}

		log.Println("Grpc Server Running")

	}()

	wg.Add(1)

	go func(config config.Config) {
		log.Println("Starting Events Server")

		defer wg.Done()
		err := events.CreateTopicIfNotExists(&config)

		if err != nil {
			log.Println(err)
			return
		}
		events.ConsumeUserCreationTopic(&config)

	}(*konfig)

	wg.Wait()
}
