package main

import (
	"di-comparison/ioc/do/builders"
	"di-comparison/shared"

	"github.com/samber/do"
)

func main() {
	injector := do.New()

	do.Provide(injector, builders.BuildParam)

	do.ProvideNamed(injector, "first-dep", builders.BuildDependency)
	do.ProvideNamed(injector, "second-dep", builders.BuildDependency)

	do.Provide(injector, builders.BuildApplication)

	app := do.MustInvoke[*shared.Application](injector)

	app.Run()
}
