package daemon

import "github.com/docker/docker/client"

func NewImageHandler(client *client.Client) *ImageHandler {
	return &ImageHandler{docker: client}
}

type ImageHandler struct {
	docker *client.Client
}
