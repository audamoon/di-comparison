package builders

import (
	"di-comparison/shared"
	"fmt"

	"go.uber.org/dig"
)

func BuildApplicationDig(container *dig.Container) (*shared.Application, error) {
	var appInstance *shared.Application

	if err := container.Invoke(func(param struct {
		dig.In
		FirstDep  *shared.SomeDependency `name:"the-first"`
		SecondDep *shared.SomeDependency `name:"the-second"`
	}) {
		appInstance = shared.NewApplication(param.FirstDep, param.SecondDep)
	}); err != nil {
		return nil, fmt.Errorf("provide: %w", err)
	}

	return appInstance, nil
}
