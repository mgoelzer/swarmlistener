package main

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	client *dockerclient.DockerClient
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/services" {
		h.services(w, r)
		return
	} else {
		io.WriteString(w, "No handler for:  "+r.URL.String())
	}
}

func (h *Handler) services(w http.ResponseWriter, r *http.Request) {
	log.Println("services() callback:  /services")
}

func HttpListen(port int, client *dockerclient.DockerClient) {
	handler := &Handler{}
	handler.client = client
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	server.ListenAndServe() // blocks
}
