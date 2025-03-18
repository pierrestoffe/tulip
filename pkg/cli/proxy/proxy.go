// Package proxy implements the proxy-related commands for managing the Tulip proxy server
package proxy

import (
	"github.com/spf13/cobra"
)

// Cmd represents the base proxy command
var Cmd = &cobra.Command{
	Use:   "proxy",
	Short: "Manage the Tulip proxy server",
	Long:  `Commands for starting, stopping, and restarting the Tulip proxy server.`,
}
