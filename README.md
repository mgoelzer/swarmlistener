# swarmlistener

First set an environment variable tha that points to your Swarm master:
```
export DOCKER_HOST=tcp://192.168.33.11:3375     # point at swarm master
```

Then run:

```
go run swarmlistener.go endpoints.go event.go
```

Current state:
* Starts a couple of nginx containers on the Swarm master using `samalba/dockerclient`
* Listens on port 3377 for HTTP requests (current none of the endpoints do anything)
* Listens to the Swarm master and logs events to the console
