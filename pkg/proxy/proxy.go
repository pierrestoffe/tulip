// Package proxy provides functionality for managing the Tulip proxy service
package proxy

import (
	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/proxy/container"
	"github.com/pierrestoffe/tulip/pkg/proxy/network"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Start initializes and launches both the proxy network and container
// Returns an error if either component fails to start
func Start() error {
	successNetwork, err := network.Start()
	if err != nil {
		return err
	}
	successContainer, err := container.Start()
	if err != nil {
		return err
	}

	if successNetwork || successContainer {
		util.PrintSuccess("Tulip's proxy was successfully started!")
	}

	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return util.HandleError("Failed to load configuration", err)
	}
	util.PrintSuccess("Access the dashboard: http://localhost:" + cfg.Proxy.AdminPort)
	return nil
}

// Stop terminates both the proxy container and network
// Returns an error if either component fails to stop
func Stop() error {
	successContainer, err := container.Stop()
	if err != nil {
		return err
	}
	successNetwork, err := network.Stop()
	if err != nil {
		return err
	}

	if successNetwork || successContainer {
		util.PrintSuccess("Tulip's proxy was successfully stopped")
	}
	return nil
}

// Restart performs a clean shutdown and restart of the proxy service
// Returns an error if either the stop or start operations fail
func Restart() error {
	if err := Stop(); err != nil {
		return err
	}
	return Start()
}

// Ensure verifies that both the network and container are running
// Starts them if they are not already running
func Ensure() error {
	if err := network.Ensure(); err != nil {
		return err
	}
	if err := container.Ensure(); err != nil {
		return err
	}
	return nil
}
