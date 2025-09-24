package builders

import (
	"di-comparison/shared"

	"github.com/samber/do"
)

func BuildParam(_ *do.Injector) (shared.Param, error) {
	return shared.NewParam(), nil
}

func BuildDependency(i *do.Injector) (*shared.SomeDependency, error) {
	return shared.NewSomeDependency(
		do.MustInvoke[shared.Param](i),
	), nil
}

func BuildApplication(i *do.Injector) (*shared.Application, error) {
	return shared.NewApplication(
		do.MustInvokeNamed[*shared.SomeDependency](i, "first-dep"),
		do.MustInvokeNamed[*shared.SomeDependency](i, "second-dep"),
	), nil
}
