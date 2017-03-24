package main

import (
	"log"

	"github.com/dnab-io/dnab/pkg/runtime/docker"
	"github.com/dnab-io/dnab/pkg/types"
)

func main() {
	engine := docker.NewDockerEngine()

	spec := types.NewContainerSpec()
	spec.Name = "dnab_example"
	spec.Image = "alpine:3.5"
	spec.Command = "/bin/sh -c \"echo $DNAB_SCRIPT | base64 -d > /tmp/ci.sh; chmod +x /tmp/ci.sh; /tmp/ci.sh\""
	spec.Env["DNAB_SCRIPT"] = "ZWNobyAiSGVsbG8sIFdvcmxkISI="

	container, err := engine.Create(spec)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(container)
}
