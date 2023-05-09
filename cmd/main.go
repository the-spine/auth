package main

import (
	"auth/internal/config"
	"fmt"
)

var konfig *config.Config

func init() {
	var err error
	konfig, err = config.LoadConfig("./")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(*konfig)
}
