package main

import "github.com/akuera/go-template/cmd/server"

func main() {
	cfg := server.EnvConfig()
	server.RunServer(cfg)
}
