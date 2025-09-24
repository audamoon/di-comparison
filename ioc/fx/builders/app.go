package builders

import (
	"context"
	"di-comparison/shared"

	"go.uber.org/fx"
)

func BuildFxApplication() *fx.App {
	return fx.New(
		fx.Provide(
			shared.NewParam,
			fx.Annotated{
				Name:   "the-first",
				Target: shared.NewSomeDependency,
			},
			fx.Annotated{
				Name:   "the-second",
				Target: shared.NewSomeDependency,
			},
		),
		fx.Invoke(StartApplication),
	)
}

type FxParams struct {
	fx.In
	FirstDep  *shared.SomeDependency `name:"the-first"`
	SecondDep *shared.SomeDependency `name:"the-second"`
}

func StartApplication(lc fx.Lifecycle, params FxParams) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			params.FirstDep.Foo()
			params.SecondDep.Foo()

			return nil
		},
		OnStop: func(context.Context) error {
			params.FirstDep.Boo()
			params.SecondDep.Boo()

			return nil
		},
	})

	return nil
}
