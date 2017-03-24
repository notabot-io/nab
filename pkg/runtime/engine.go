package runtime

import (
	"io"

	"github.com/dnab-io/dnab/pkg/types"
)

// Engine defines the interface for interacting with the container-based backend.
type Engine interface {
	// Create takes a specification for a container, and attempts to create (but not run) a
	// container that matches the specification.
	Create(*types.ContainerSpec) (*types.Container, error)
	// Run takes a container produced by Create and runs it.
	Run(*types.Container) error
	// Tail outputs the logs from the given container.
	Tail(*types.Container) (io.ReadCloser, error)
	// Wait for a container to exit on it's own.
	Wait(*types.Container) (*types.ContainerState, error)
	// Stop attempts to gracefully stop a given container.
	Stop(*types.Container) error
	// Kill attempts to instantly stop a given container.
	Kill(*types.Container) error
}
