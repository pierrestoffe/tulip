// Package network provides functionality for managing the Docker network used by Tulip's proxy
package network

import (
	"bytes"
	"os/exec"

	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Start creates and initializes a new Docker network for the proxy if one doesn't exist
// Returns true if a new network was created, false if it already existed, and any error that occurred
func Start() (bool, error) {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return false, util.HandleError("Failed to load configuration", err)
	}

	// Check if the proxy network is already running
	if IsRunning() {
		util.PrintWarning("Proxy network " + cfg.Docker.NetworkName + " is already running")
		return false, nil
	}

	// Start the proxy network
	util.PrintInfo("Starting " + cfg.Docker.NetworkName + " proxy network..")
	cmd := exec.Command("docker", "network", "create", cfg.Docker.NetworkName)

	// Capture stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run command and handle errors
	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		return false, util.HandleError("Error starting "+cfg.Docker.NetworkName+" proxy network", err, errMsg)
	}

	util.PrintInfoReplace("Proxy network " + cfg.Docker.NetworkName + " started")
	return true, nil
}

// Stop removes the Docker proxy network
// Returns true if the network was stopped, false if it wasn't running, and any error that occurred
func Stop() (bool, error) {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return false, util.HandleError("Failed to load configuration", err)
	}

	// Check if the proxy network is running
	if !IsRunning() {
		util.PrintWarning("Proxy network " + cfg.Docker.NetworkName + " is already stopped.")
		return false, nil
	}

	util.PrintInfo("Stopping " + cfg.Docker.NetworkName + " network..")

	// Stop the proxy network
	cmd := exec.Command("docker", "network", "remove", cfg.Docker.NetworkName)

	// Capture stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run command and handle errors
	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		return false, util.HandleError("Error stopping "+cfg.Docker.NetworkName+" proxy network", err, errMsg)
	}

	util.PrintInfoReplace("Proxy network " + cfg.Docker.NetworkName + " was stopped")
	return true, nil
}

// Ensure verifies that the proxy network exists and creates it if missing
// Returns an error if the network cannot be created
func Ensure() error {
	if IsRunning() {
		return nil
	}
	_, err := Start()
	return err
}

// IsRunning checks if the proxy network exists in Docker
// Returns true if the network exists and is operational
func IsRunning() bool {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		util.HandleError("Failed to load configuration", err)
		return false
	}

	cmd := exec.Command("docker", "network", "inspect", cfg.Docker.NetworkName)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
