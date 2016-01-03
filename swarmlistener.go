package main

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	client *dockerclient.DockerClient
)

func main() {
	// Example input data:
	serviceName := "nginx"
	serviceImage := "nginx"
	//serviceCmd := []string{""}
	serviceCnt := 2

	// DOCKER_HOST should point to Swarm master, e.g., "tcp://192.168.33.11:3375"
	docker, err := dockerclient.NewDockerClient(os.Getenv("DOCKER_HOST"), nil)
	if err != nil {
		log.Fatal(err)
	}

	client = docker

	client.StartMonitorEvents(EventCallback, nil)

	// Start Service
	containerIds := make([]string, serviceCnt)
	for i := 0; i < serviceCnt; i++ {
		containerConfig := &dockerclient.ContainerConfig{
			Image: serviceImage,
			//Cmd:   serviceCmd,
		}
		containerName := fmt.Sprintf("%s-%02d", serviceName, i)
		fmt.Println(containerName)
		_ = containerConfig
		containerId, err := docker.CreateContainer(containerConfig, containerName, nil)
		if err != nil {
			log.Fatal(err)
		}
		hostConfig := &dockerclient.HostConfig{}
		err = docker.StartContainer(containerId, hostConfig)
		if err != nil {
			log.Fatal(err)
		}
		containerIds[i] = containerId
	}
	for _, containerId := range containerIds {
		fmt.Printf("Started %v\n", containerId)
	}

	// Block until process is killed
	sigChan := make(chan os.Signal, 3)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for _ = range sigChan {
			// Cleanup
			fmt.Printf("Cleaning up\n")
			client.StopAllMonitorEvents()
			os.Exit(0)
		}
	}()
	HttpListen(3377, client) // blocks
}
