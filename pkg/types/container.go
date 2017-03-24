package types

// Container helps to identify a created container.
type Container struct {
	// ID is a unique identifier for the container.
	ID string
	// Spec is the ContainerSpec that was used to create this container.
	Spec *ContainerSpec
}
