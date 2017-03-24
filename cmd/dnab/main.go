package main

import (
	"fmt"
	"log"

	"github.com/dnab-io/dnab/pkg/runtime/docker"
	"github.com/dnab-io/dnab/pkg/types"
)

func main() {
	engine := docker.NewDockerEngine()

	spec := types.NewContainerSpec()
	spec.Name = "dnab_example"
	spec.Image = "alpine:3.5"
	spec.Env["DNAB_SCRIPT"] = "ZWNobyAiSGVsbG8sIFdvcmxkISI="

	scriptName := "/tmp/ci.sh"

	spec.EntryPoint = []string{"/bin/sh", "-c"}
	spec.Command = []string{fmt.Sprintf(
		"echo $DNAB_SCRIPT | base64 -d > %s; chmod +x %s; %s; sleep 120",
		scriptName,
		scriptName,
		scriptName,
	)}

	container, err := engine.Create(spec)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(container.ID)
}
