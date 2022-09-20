package dockerHost

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImages() []string {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unable to get new docker client: %v", err)
		return nil
	}
	containers := getContainers(dockerClient)
	var images []string
	for _, container := range containers {
		images = append(images, container.Image)
	}
	return images
}

func getContainers(cli *client.Client) []types.Container {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Printf("Unable to list containers: %v", err)
		return nil
	}
	return containers
}
