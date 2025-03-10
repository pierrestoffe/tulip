package proxy

import (
    "github.com/pierrestoffe/tulip/pkg/proxy"
    "github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
    Use:   "restart",
    Short: "Restart the Tulip proxy server",
    Long:  `Restart the Tulip proxy server.`,
    Run: func(cmd *cobra.Command, args []string) {
        proxy.Restart()
    },
}

func init() {
    Cmd.AddCommand(RestartCmd)
}
