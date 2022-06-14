package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-jackdemo/demo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: demo.Provider,
	})
}
