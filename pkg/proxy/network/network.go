package container

import (
    "os/exec"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/util/print"
)

// Checks if the Docker proxy network exists
// Returns true if the network exists, false otherwise
func isRunning() bool {
    cmd := exec.Command("docker", "network", "inspect", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        return false
    }
    return true
}

// Creates the Docker proxy network if it doesn't already exist
// Returns nil if the network already exists or was created successfully
func Start() error {
    // Check if the proxy network is already running
    if isRunning() {
        print.Warning("Proxy network " + constants.ProxyNetworkName + " is running already.")
        return nil
    }

    // Start the proxy network
    print.Info("Starting " + constants.ProxyNetworkName + " proxy network...")
    cmd := exec.Command("docker", "network", "create", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        return helpers.HandleError("error starting " + constants.ProxyNetworkName + " proxy network", err)
    }
    print.SuccessReplace("Proxy network " + constants.ProxyNetworkName + " started.")
    return nil
}

// Removes the Docker proxy network if it exists
// Returns nil if the network doesn't exist or was removed successfully
func Stop() error {
    // Check if the proxy network is running
    if !isRunning() {
        print.Warning("Proxy network " + constants.ProxyNetworkName + " was stopped already.")
        return nil
    }

    print.Info("Stopping " + constants.ProxyNetworkName + " network...")

    // Stop the proxy network
    cmd := exec.Command("docker", "network", "remove", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        return helpers.HandleError("error stopping " + constants.ProxyNetworkName + " proxy network", err)
    }
    print.SuccessReplace("Proxy network " + constants.ProxyNetworkName + " was stopped.")
    return nil
}
