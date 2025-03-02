package main

import "github.com/alecthomas/kong"

var args struct {
	Version        kong.VersionFlag `name:"version" short:"v" help:"Show version information."`
	ConfigFilePath string           `name:"config" short:"c" optional:"" type:"path" default:"./config.yaml" help:"Specify path to configuration file"`
}
