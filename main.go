package main

import (
	"github.com/NubeIO/module-core-loraraw/logger"
	"github.com/NubeIO/module-core-loraraw/pkg"
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/hashicorp/go-plugin"
)

func ServePlugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins: plugin.PluginSet{
			"nube-module": &shared.NubeModule{Impl: &pkg.Module{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

func main() {
	logger.SetLogger("INFO")
	ServePlugin()
}
