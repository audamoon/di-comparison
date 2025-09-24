package main

import (
	"di-comparison/ioc/sarulabs/builders"
	"log"
)

func main() {
	container, err := builders.BuildSarulabsContainer()
	if err != nil {
		log.Fatal(err)
	}

	app, err := builders.BuildApplicationSarulabs(container)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
