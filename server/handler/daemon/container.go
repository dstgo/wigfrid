package daemon

import "github.com/docker/docker/client"

func NewContainerHandler(client *client.Client) *ContainerHandler {
	return &ContainerHandler{docker: client}
}

type ContainerHandler struct {
	docker *client.Client
}
