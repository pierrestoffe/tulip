// Package start implements the 'start' command functionality
package start

import (
	"github.com/pierrestoffe/tulip/pkg/project"
	"github.com/pierrestoffe/tulip/pkg/proxy"
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/spf13/cobra"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "Start the project",
	Long:  `Start the project that is found in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure Tulip is properly set up
		if err := setup.Ensure(); err != nil {
			return
		}
		// Ensure proxy service is running
		if err := proxy.Ensure(); err != nil {
			return
		}

		project.Start()
	},
}
