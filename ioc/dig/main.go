package main

import (
	"di-comparison/ioc/dig/builders"
	"di-comparison/shared"
	"log"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	if err := container.Provide(shared.NewParam); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(shared.NewSomeDependency, dig.Name("the-first")); err != nil {
		log.Fatal(err)
	}

	if err := container.Provide(shared.NewSomeDependency, dig.Name("the-second")); err != nil {
		log.Fatal(err)
	}

	app, err := builders.BuildApplicationDig(container)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
