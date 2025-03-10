package proxy

import (
    "github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
    Use:   "proxy",
    Short: "Manage the Tulip proxy server",
    Long:  `Commands for starting, stopping, and restarting the Tulip proxy server.`,
}

func init() {
}
