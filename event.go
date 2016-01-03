package main

import (
	"github.com/samalba/dockerclient"
	"log"
)

func EventCallback(e *dockerclient.Event, ec chan error, args ...interface{}) {
	log.Printf(">> Event:\n")
	log.Printf(">>     Container Id:  %v):\n", e.Id)
	log.Printf(">>     Status:        %v\n", e.Status)
	log.Printf(">>     From:          '%v'\n", e.From)
	log.Printf(">>     Time:          %v\n", e.Time)
}
