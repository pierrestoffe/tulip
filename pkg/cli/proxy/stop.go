// Package proxy implements the proxy command functionality
package proxy

import (
	"github.com/pierrestoffe/tulip/pkg/proxy"
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/spf13/cobra"
)

// StopCmd represents the proxy stop command
// It ensures proper setup and stops the proxy service
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Tulip proxy server",
	Long:  `Stop the Tulip proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setup.Ensure(); err != nil {
			return
		}

		proxy.Stop()
	},
}

func init() {
	Cmd.AddCommand(StopCmd)
}
