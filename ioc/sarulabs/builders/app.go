package builders

import (
	"di-comparison/shared"

	"github.com/sarulabs/di/v2"
)

func BuildSarulabsContainer() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return di.Container{}, err
	}

	// Register Param
	err = builder.Add(di.Def{
		Name: "param",
		Build: func(ctn di.Container) (interface{}, error) {
			return shared.NewParam(), nil
		},
	})
	if err != nil {
		return di.Container{}, err
	}

	// Register first dependency
	err = builder.Add(di.Def{
		Name: "first-dep",
		Build: func(ctn di.Container) (interface{}, error) {
			param := ctn.Get("param").(shared.Param)
			return shared.NewSomeDependency(param), nil
		},
	})
	if err != nil {
		return di.Container{}, err
	}

	// Register second dependency
	err = builder.Add(di.Def{
		Name: "second-dep",
		Build: func(ctn di.Container) (interface{}, error) {
			param := ctn.Get("param").(shared.Param)
			return shared.NewSomeDependency(param), nil
		},
	})
	if err != nil {
		return di.Container{}, err
	}

	return builder.Build(), nil
}

func BuildApplicationSarulabs(container di.Container) (*shared.Application, error) {
	firstDep := container.Get("first-dep").(*shared.SomeDependency)
	secondDep := container.Get("second-dep").(*shared.SomeDependency)

	return shared.NewApplication(firstDep, secondDep), nil
}
