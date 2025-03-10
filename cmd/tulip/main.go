package main

import (
    "os"
	"github.com/pierrestoffe/tulip/pkg/cli"
)

func main() {
    if err := cli.Execute(); err != nil {
        os.Exit(1)
    }
}
