package server

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port  int `required:"true" default:"6100"`
	Mongo MongoConfig
}

type MongoConfig struct {
	Host     string `required:"true"`
	Scheme   string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
}

func EnvConfig() *Config {
	var cfg Config
	err := envconfig.Process("DOM", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
