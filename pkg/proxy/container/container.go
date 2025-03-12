package container

import (
    "bytes"
    "os"
    "os/exec"
    "strings"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/util/print"
)

// Checks if the proxy container is currently running
// Returns true if the container is running, false otherwise
func isRunning() bool {
    cmd := exec.Command("docker", "ps", "--filter", "name=" + constants.ProxyContainerName, "--format", "{{.Names}}")
    output, err := cmd.Output()
    if err != nil {
        return false
    }
    return len(strings.TrimSpace(string(output))) > 0
}

// Launches the proxy container if it's not already running
// Returns nil if the container is already running or was started successfully
func Start() error {
    // Check if the proxy container is already running
    if isRunning() {
        print.Warning("Proxy " + constants.ProxyContainerName + " is running already.")
        print.Info("You can access the dashboard at " + constants.ProxyUrl)
        return nil
    }

    print.Info("Starting " + constants.ProxyContainerName + " proxy...")

    // Get the path to the proxy configuration directory
    proxyConfigDir, err := helpers.GetProxyConfigDir()
    if err != nil {
        return err
    }

    // Start the proxy container
    cmd := exec.Command("docker", "compose", "up", "-d")
    cmd.Dir = proxyConfigDir
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "COMPOSE_IGNORE_ORPHANS=1")

    // Capture stderr
    var stderr bytes.Buffer
    cmd.Stderr = &stderr

    // Run command and handle errors
    if err := cmd.Run(); err != nil {
        errMsg := stderr.String()
        return helpers.HandleError("error starting " + constants.ProxyContainerName + " proxy: ", err, errMsg)
    }

    print.SuccessReplace("Proxy " + constants.ProxyContainerName + " started.")
    print.Info("You can access the dashboard at " + constants.ProxyUrl)
    return nil
}

// Terminates the proxy container if it's running
// Returns nil if the container isn't running or was stopped successfully
func Stop() error {
    // Check if the proxy container is running
    if !isRunning() {
        print.Warning("Proxy " + constants.ProxyContainerName + " was stopped already.")
        return nil
    }

    print.Info("Stopping " + constants.ProxyContainerName + " proxy...")

    // Get the path to the proxy configuration directory
    proxyConfigDir, err := helpers.GetProxyConfigDir()
    if err != nil {
        return err
    }

    // Stop the proxy container
    cmd := exec.Command("docker", "compose", "down")
    cmd.Dir = proxyConfigDir
    cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "COMPOSE_IGNORE_ORPHANS=1")

    // Capture stderr
    var stderr bytes.Buffer
    cmd.Stderr = &stderr

    // Run command and handle errors
    if err := cmd.Run(); err != nil {
        errMsg := stderr.String()
        return helpers.HandleError("error stopping " + constants.ProxyContainerName + " proxy: ", err, errMsg)
    }

    print.SuccessReplace("Proxy " + constants.ProxyContainerName + " was stopped.")
    return nil
}

// Ensure checks if the proxy container is running and starts it if it's not
// Returns nil if the container is already running or was started successfully
func Ensure() error {
    if isRunning() {
        return nil
    }
    return Start()
}
