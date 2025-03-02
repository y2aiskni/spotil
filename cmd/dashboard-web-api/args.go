package main

import "github.com/alecthomas/kong"

var args struct {
	Version kong.VersionFlag `name:"version" short:"v" help:"Show version information."`
}
