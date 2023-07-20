package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

var conf Config

func InitConfig() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}

func GetConfig() Config {
	return conf
}
