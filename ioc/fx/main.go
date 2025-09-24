package main

import (
	"context"
	"di-comparison/ioc/fx/builders"
	"log"
)

func main() {
	app := builders.BuildFxApplication()

	app.Run()

	if err := app.Stop(context.TODO()); err != nil {
		log.Println(err)
	}
}
