package docker

import (
	"io"

	"github.com/dnab-io/dnab/pkg/runtime"
	"github.com/dnab-io/dnab/pkg/types"
)

type dockerEngine struct {
}

// NewDockerEngine returns a new runtime.Engine instance powered by Docker.
func NewDockerEngine() runtime.Engine {
	return &dockerEngine{}
}

func (e *dockerEngine) Create(*types.ContainerSpec) (*types.Container, error) {
	container := &types.Container{}

	return container, nil
}

func (e *dockerEngine) Run(*types.Container) error {
	return nil
}

func (e *dockerEngine) Tail(*types.Container) (io.ReadCloser, error) {
	return nil, nil
}

func (e *dockerEngine) Wait(*types.Container) (*types.ContainerState, error) {
	state := &types.ContainerState{}

	return state, nil
}

func (e *dockerEngine) Stop(*types.Container) error {
	return nil
}

func (e *dockerEngine) Kill(*types.Container) error {
	return nil
}
