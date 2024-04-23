package daemon

import "github.com/docker/docker/client"

func NewDaemonHandler(client *client.Client) *DaemonHandler {
	return &DaemonHandler{docker: client}
}

type DaemonHandler struct {
	docker *client.Client
}
