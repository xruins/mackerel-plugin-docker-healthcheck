package main

import (
	"log"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/k0kubun/pp"
)

func main() {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		panic(err)
	}

	ids := make([]string, 0, len(containers))
	for _, c := range containers {
		ids = append(ids, c.ID)
	}

	// state := make(map[string]string, len(containers))
	for _, id := range ids {
		container, err := client.InspectContainer(id)
		if err != nil {
			log.Panicf("failed to get detailed information of the  container. id: %s, err: %s", id, err)
		}

		pp.Print(container)
	}
}
