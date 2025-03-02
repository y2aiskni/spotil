package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var version = "0.0.0"

func main() {
	ctx := kong.Parse(&args, &kong.Vars{"version": version})
	switch ctx.Command() {
	case "":
		fmt.Println("execute!")
	default:
		panic(ctx.Command())
	}
}
