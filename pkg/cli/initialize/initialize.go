// Package initialize implements the 'init' command functionality for setting up Tulip
package initialize

import (
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/spf13/cobra"
)

// Cmd represents the initialization command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Tulip",
	Long:  `Initialize Tulip by creating missing directories and config files in the home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		setup.Initialize()
	},
}
