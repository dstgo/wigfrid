package types

import "github.com/docker/docker/client"

type Env struct {
	// docker http api client
	Docker *client.Client
}
