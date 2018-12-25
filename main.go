package main

import (
	"./api"
	"./conf"
	"./mysql"
	"./server"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(conf.NewConfig)
	return container
}

func main() {
	container := BuildContainer()
	container.Provide(conf.NewConfig)
	container.Provide(mysql.ConnectDatabase)
	container.Provide(mysql.NewPersonRepository)
	container.Provide(api.NewPersonService)
	container.Provide(server.NewServer)

	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}
