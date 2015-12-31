package main

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func eventCallback(e *dockerclient.Event, ec chan error, args ...interface{}) {
	log.Printf(">> Event (Id=%v):\n", e.Id)
	log.Printf(">>     Status:  %v\n", e.Status)
	log.Printf(">>     From:    '%v'\n", e.From)
	log.Printf(">>     Time:    %v\n", e.Time)
}

var (
	client *dockerclient.DockerClient
)

func waitForInterrupt() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	for _ = range sigChan {
		client.StopAllMonitorEvents()
		os.Exit(0)
	}
}

func main() {
	// Example input data:
	serviceName := "nginx"
	serviceImage := "nginx"
	serviceCmd := []string{""}
	serviceCnt := 3

	// DOCKER_HOST should point to Swarm master, e.g., "tcp://192.168.33.11:3375"
	docker, err := dockerclient.NewDockerClient(os.Getenv("DOCKER_HOST"), nil)
	if err != nil {
		log.Fatal(err)
	}

	client = docker

	client.StartMonitorEvents(eventCallback, nil)

	// Start Service
	containerIds := make([]string, serviceCnt)
	for i := 0; i < serviceCnt; i++ {
		containerConfig := &dockerclient.ContainerConfig{
			Image: serviceImage,
			Cmd:   serviceCmd,
		}
		containerName := fmt.Sprintf("%s-%02d", serviceName, i)
		fmt.Println(containerName)
		_ = containerConfig
		containerId, err := docker.CreateContainer(containerConfig, containerName, nil)
		if err != nil {
			log.Fatal(err)
		}
		containerIds[i] = containerId
	}
	for _,containerId := range containerIds {
		fmt.Printf("Started %v\n",containerId)
	}

	defer func() {
		// Cleanup Service
		fmt.Printf("Cleaning up...\n")
		// TODO
	}()

	waitForInterrupt()
}
