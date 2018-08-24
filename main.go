package main

import (
	"github.com/crawford/terraform-provider-libvirt/libvirt"
	"github.com/hashicorp/terraform/plugin"
	"math/rand"
	"time"
)

func main() {
	defer libvirt.CleanupLibvirtConnections()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: libvirt.Provider,
	})
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
