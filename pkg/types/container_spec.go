package types

// @todo: Is ContainerSpec just going to be the same as Container in the end? Perhaps not, Container
// on it's own may have some information to identify a container, and have a spec attached?

// ContainerSpec defines the desired state of a container to be created.
type ContainerSpec struct {
	// EntryPoint is the executable to run. The Command will be passed as an argument.
	EntryPoint []string
	// Command provides the 'script' that will be run when starting the container.
	Command []string
	// Dns is a list of DNS server addresses (<host>[:<port>]).
	Dns []string
	// Env is a list of environment variables to pass to the container.
	Env map[string]string
	// Image provides the name of an image to run.
	Image string
	// Name is the name the container will use once created.
	Name string
	// User is the user the commands in the container will run as (<name|uid>[:<group|gid>]).
	User string
	// Volumes are a list of volumes to mount.
	Volumes []string
	// WorkDir specifies the working directory that commands will run relative to.
	WorkDir string
}

// NewContainerSpec creates a new ContainerSpec with fields initialised.
func NewContainerSpec() *ContainerSpec {
	spec := &ContainerSpec{}
	spec.Env = make(map[string]string)

	return spec
}
