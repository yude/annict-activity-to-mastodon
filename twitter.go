package main

import (
	"log"

	"github.com/sivchari/gotwtr"
)

func GetTwitterClient() *gotwtr.Client {
	// Parse config.toml
	conf, err := NewConfig()
	if err != nil {
		log.Fatal("Error: Failed to parse config.toml")
	}

	client := gotwtr.New(conf.Credentials.TwitterBearer)

	return client
}
