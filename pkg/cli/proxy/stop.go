package proxy

import (
    "github.com/pierrestoffe/tulip/pkg/proxy"
    "github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
    Use:   "stop",
    Short: "Stop the Tulip proxy server",
    Long:  `Stop the Tulip proxy server.`,
    Run: func(cmd *cobra.Command, args []string) {
        proxy.Stop()
    },
}

func init() {
    Cmd.AddCommand(StopCmd)
}
