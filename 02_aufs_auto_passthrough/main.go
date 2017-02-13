package main

import (
	"github.com/atlantic777/docker_graphdriver_plugins/aufs"
	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/go-plugins-helpers/graphdriver/shim"
)

func main() {
	if reexec.Init() {
		return
	}

	h := shim.NewHandlerFromGraphDriver(aufs.Init)
	h.ServeUnix("plugin", 0)
}
