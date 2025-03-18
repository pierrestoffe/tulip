// Package proxy implements the proxy command functionality
package proxy

import (
	"github.com/pierrestoffe/tulip/pkg/proxy"
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/spf13/cobra"
)

// StartCmd represents the proxy start command
// It ensures proper setup and starts the proxy service
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Tulip proxy server",
	Long:  `Start the Tulip proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setup.Ensure(); err != nil {
			return
		}

		proxy.Start()
	},
}

func init() {
	Cmd.AddCommand(StartCmd)
}
