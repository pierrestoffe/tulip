package initialize

import (
    "github.com/pierrestoffe/tulip/pkg/app"
    "github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize Tulip",
    Long:  `Initialize Tulip by creating missing directories and config files in the home directory.`,
    Run: func(cmd *cobra.Command, args []string) {
        app.Initialize()
    },
}

func init() {
}
