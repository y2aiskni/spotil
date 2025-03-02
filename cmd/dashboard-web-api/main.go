package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&args)
	switch ctx.Command() {
	case "":
		fmt.Println("execute!")
	default:
		panic(ctx.Command())
	}
}
