// Package main is the entry point for the Tulip command-line application
package main

import (
	"os"

	"github.com/pierrestoffe/tulip/pkg/cli"
)

// main initializes and executes the Tulip CLI application
func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
