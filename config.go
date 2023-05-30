package main

import (
	"github.com/BurntSushi/toml"
)

var conf Config

func NewConfig() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
