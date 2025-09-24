package builders

import (
	"di-comparison/shared"

	"github.com/golobby/container/v3"
)

func NewGoLobbyApplication(c container.Container) (*shared.Application, error) {
	var dep1, dep2 *shared.SomeDependency

	if err := c.NamedResolve(&dep1, "first-dep"); err != nil {
		return nil, err
	}

	if err := c.NamedResolve(&dep2, "second-dep"); err != nil {
		return nil, err
	}

	return shared.NewApplication(dep1, dep2), nil
}
