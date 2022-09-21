package dockerHost

import (
	"bytes"
	"context"
	"log"
	"strings"

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

func GetLsForContainer(containerName string) string {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Unable to get new docker client: %v", err)
		return ""
	}
	containers := getContainers(dockerClient)
	for _, container := range containers {
		log.Println("Container: ", container.Names, " Image: ", container.Image)
		if strings.Contains(container.Names[0], containerName) {
			log.Println("Container found")
			return getExecResultForContainer(dockerClient, container.ID, []string{"ls", "-l"})
		}
	}
	log.Fatalf("Container not found: %v", containerName)
	return ""
}

func getExecResultForContainer(cli *client.Client, containerId string, cmd []string) (execResult string) {
	var execConfig types.ExecConfig = types.ExecConfig{
		Cmd:          cmd,
		AttachStdout: true,
		AttachStderr: true,
	}
	execId, err := cli.ContainerExecCreate(context.Background(), containerId, execConfig)
	if err != nil {
		log.Printf("Unable to create exec: %v", err)
		return ""
	}
	res, _ := cli.ContainerExecAttach(context.Background(), execId.ID, types.ExecStartCheck{})
	defer res.Close()
	err = cli.ContainerExecStart(context.Background(), execId.ID, types.ExecStartCheck{})
	if err != nil {
		log.Printf("Unable to start exec: %v", err)
		return ""
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Reader)
	if err != nil {
		log.Printf("Unable to read from exec: %v", err)
		return ""
	}
	execResult = buf.String()
	return
}
