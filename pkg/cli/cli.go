// Package cli handles the command-line interface functionality of the Tulip application
package cli

import (
	"github.com/pierrestoffe/tulip/pkg/cli/initialize"
	"github.com/pierrestoffe/tulip/pkg/cli/proxy"
	"github.com/pierrestoffe/tulip/pkg/cli/start"
	"github.com/pierrestoffe/tulip/pkg/setup"
	"github.com/pierrestoffe/tulip/pkg/util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command for the Tulip CLI application.
// It serves as the entry point for all Tulip commands and subcommands.
var rootCmd = &cobra.Command{
	Use:   "tulip",
	Short: "Tulip description",
	Long:  `Long Tulip description`,
}

// Execute runs the root command and handles any errors that occur
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return util.PrintErrorE(err)
	}
	return nil
}

// ValidateSetup checks if Tulip is properly set up before running commands
// It skips validation for the "init" command since setup isn't required for initialization
func ValidateSetup(cmdName string) error {
	if cmdName == "init" {
		return nil
	}
	return setup.Ensure()
}

// init adds all child commands to the root command
func init() {
	rootCmd.AddCommand(proxy.Cmd)
	rootCmd.AddCommand(initialize.Cmd)
	rootCmd.AddCommand(start.Cmd)
}
