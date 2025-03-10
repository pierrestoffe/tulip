package cli

import (
    "github.com/pierrestoffe/tulip/pkg/cli/proxy"
    "github.com/pierrestoffe/tulip/pkg/cli/initialize"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "tulip",
    Short: "Tulip description",
    Long:  `Long Tulip description`,
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(proxy.Cmd)
    rootCmd.AddCommand(initialize.Cmd)
}
