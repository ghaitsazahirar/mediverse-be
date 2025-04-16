package main

import (
	"mediversebe/internal/bootstrap"

	"github.com/sev-2/raiden"
)

func main() {
	// load configuration
	config, err := raiden.LoadConfig(nil)
	if err != nil {
		raiden.Error("load configuration", err.Error())
	}

	// Setup server
	server := raiden.NewServer(config)

	// register resource
	bootstrap.RegisterRoute(server)
	bootstrap.RegisterJobs(server)
	bootstrap.RegisterSubscribers(server)

	// run server
	server.Run()
}
