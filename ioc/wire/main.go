package main

import (
	"di-comparison/ioc/wire/builders"
	"log"
)

func main() {
	app, err := builders.BuildWireApplication()
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
