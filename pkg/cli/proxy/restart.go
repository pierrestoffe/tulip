// Package proxy implements the proxy command functionality
package proxy

import (
	"github.com/pierrestoffe/tulip/pkg/proxy"
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/spf13/cobra"
)

// RestartCmd represents the proxy restart command
// It ensures proper setup and restarts the proxy service
var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Tulip proxy server",
	Long:  `Restart the Tulip proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setup.Ensure(); err != nil {
			return
		}

		proxy.Restart()
	},
}

func init() {
	Cmd.AddCommand(RestartCmd)
}
