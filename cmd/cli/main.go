package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/itsahedge/wrapped-celestia/cmd/pkg"
	"os"
)

func main() {
	ctx := NewCLI()
	if err := ctx.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

var CLI struct {
	Bridged GetBridged `cmd:"" help:"GetBridgedTo Events"`
}

func NewCLI() *kong.Context {
	ctx := kong.Parse(&CLI)
	return ctx
}

func TiaClient() (*pkg.TiaClient, error) {
	return pkg.NewClient()
}
