package main

import (
	"di-comparison/ioc/golobby/builders"
	"di-comparison/shared"
	"log"

	"github.com/golobby/container/v3"
)

func main() {
	c := container.New()

	if err := c.Singleton(shared.NewParam); err != nil {
		log.Fatal(err)
	}

	if err := c.NamedSingleton("first-dep", shared.NewSomeDependency); err != nil {
		log.Fatal(err)
	}

	if err := c.NamedSingleton("second-dep", shared.NewSomeDependency); err != nil {
		log.Fatal(err)
	}

	app, err := builders.NewGoLobbyApplication(c)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
