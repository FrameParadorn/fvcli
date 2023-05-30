package main

import (
	"git.robodev.co/newfinvest/platform/common-service/internals/container"
)

func main() {
	server, err := container.NewContainer()
	if err != nil {
		panic(err)
	}

	if err := server.MigrateDB(); err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
