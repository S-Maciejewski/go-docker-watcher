package container

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListContainer List containers running on host machine
func ListContainer() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unable to get new docker client: %v", err)
		return err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Printf("Unable to list containers: %v", err)
		return err
	}
	if len(containers) > 0 {
		for _, container := range containers {
			log.Printf("Container ID: %s", container.ID)
			log.Printf("Container Image: %s", container.Image)
		}
	} else {
		log.Println("There are no containers running")
	}
	return nil
}
