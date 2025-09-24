//go:build wireinject
// +build wireinject

package builders

import (
	"di-comparison/shared"

	"github.com/google/wire"
)

type (
	FirstDep  *shared.SomeDependency
	SecondDep *shared.SomeDependency
)

func NewFirstDep(param shared.Param) FirstDep {
	return shared.NewSomeDependency(param)
}

func NewSecondDep(param shared.Param) SecondDep {
	return shared.NewSomeDependency(param)
}

func NewApplication(first FirstDep, second SecondDep) *shared.Application {
	return shared.NewApplication(first, second)
}

func BuildWireApplication() (*shared.Application, error) {
	wire.Build(
		shared.NewParam,
		NewFirstDep,
		NewSecondDep,
		NewApplication,
	)
	return nil, nil
}
