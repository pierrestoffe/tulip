package proxy

import (
    "github.com/pierrestoffe/tulip/pkg/proxy"
    "github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the Tulip proxy server",
    Long:  `Start the Tulip proxy server.`,
    Run: func(cmd *cobra.Command, args []string) {
        proxy.Start()
    },
}

func init() {
    Cmd.AddCommand(StartCmd)
}
