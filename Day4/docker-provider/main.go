package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-docker/internal/provider"
)

var (
	version string = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New(version)}

	if debugMode {
		debugOpts := &plugin.ServeOpts{ProviderFunc: provider.New(version), ProviderAddr: "registry.terraform.io/tektutor/docker", Debug: true}
		plugin.Serve(debugOpts)
		return
	}

	plugin.Serve(opts)
}
