package main

import (
	"github.com/SolaceLabs/steampipe-plugin-solace/solace"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: solace.Plugin})
}
