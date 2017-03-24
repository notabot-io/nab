package docker

import (
	"context"
	"io"

	"github.com/dnab-io/dnab/pkg/runtime"
	"github.com/dnab-io/dnab/pkg/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// dockerEngine is runtime engine that uses Docker for container management.
type dockerEngine struct {
	// Client is a Docker client.
	client *client.Client
}

// NewDockerEngine returns a new runtime.Engine instance powered by Docker.
func NewDockerEngine() runtime.Engine {
	dkrClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return &dockerEngine{
		client: dkrClient,
	}
}

func (e *dockerEngine) Create(spec *types.ContainerSpec) (*types.Container, error) {
	ctrConfig := container.Config{}
	ctrConfig.Image = spec.Image
	ctrConfig.User = spec.User
	ctrConfig.AttachStderr = true
	ctrConfig.AttachStdout = true

	for k, v := range spec.Env {
		ctrConfig.Env = append(ctrConfig.Env, k+"="+v)
	}

	ctrConfig.WorkingDir = spec.WorkDir
	ctrConfig.Entrypoint = spec.EntryPoint
	ctrConfig.Cmd = spec.Command

	hostConfig := container.HostConfig{}
	hostConfig.DNS = spec.Dns

	response, err := e.client.ContainerCreate(
		context.Background(),
		&ctrConfig,
		&hostConfig,
		&network.NetworkingConfig{},
		spec.Name,
	)

	if err != nil {
		return nil, err
	}

	ctr := &types.Container{}
	ctr.ID = response.ID
	ctr.Spec = spec

	return ctr, nil
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
