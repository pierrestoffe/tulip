// Package container manages the lifecycle of Tulip's proxy Docker container
package container

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Launches the proxy container if it's not already running
func Start() (bool, error) {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return false, util.HandleError("Failed to load configuration", err)
	}

	// Check if the proxy container is already running
	if IsRunning() {
		util.PrintWarning("Proxy " + config.ProxyContainerName + " is already running")
		return false, nil
	}

	// Verify that all used ports are open
	if err := verifyPorts(); err != nil {
		return false, err
	}

	util.PrintInfo("Starting " + config.ProxyContainerName + " proxy..")

	// Get the path to the proxy configuration directory
	proxyConfigDir, err := config.GetProxyConfigDir()
	if err != nil {
		return false, err
	}

	// Start the proxy container
	cmd := prepareDockerComposeCmd("docker", []string{"compose", "up", "-d"}, proxyConfigDir, cfg)

	// Capture stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run command and handle errors
	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		return false, util.HandleError("Error starting "+config.ProxyContainerName+" proxy", err, errMsg)
	}

	util.PrintInfoReplace("Proxy " + config.ProxyContainerName + " started")
	return true, nil
}

// Terminates the proxy container if it's running
func Stop() (bool, error) {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return false, util.HandleError("Failed to load configuration", err)
	}

	// Check if the proxy container is running
	if !IsRunning() {
		util.PrintWarning("Proxy " + config.ProxyContainerName + " is already stopped.")
		return false, nil
	}

	util.PrintInfo("Stopping " + config.ProxyContainerName + " proxy..")

	// Get the path to the proxy configuration directory
	proxyConfigDir, err := config.GetProxyConfigDir()
	if err != nil {
		return false, err
	}

	// Stop the proxy container
	cmd := prepareDockerComposeCmd("docker", []string{"compose", "down", "--remove-orphans"}, proxyConfigDir, cfg)

	// Capture stderr
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// Run command and handle errors
	if err := cmd.Run(); err != nil {
		errMsg := stderr.String()
		return false, util.HandleError("Error stopping "+config.ProxyContainerName+" proxy", err, errMsg)
	}

	util.PrintInfoReplace("Proxy " + config.ProxyContainerName + " was stopped")
	return true, nil
}

// Ensure checks if the proxy container is running and starts it if it's not
func Ensure() error {
	if IsRunning() {
		return nil
	}
	_, err := Start()
	return err
}

// Checks if the proxy container is currently running
func IsRunning() bool {
	cmd := exec.Command("docker", "ps", "--filter", "name="+config.ProxyContainerName, "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		util.HandleError("Failed to check if container is running", err)
		return false
	}
	return len(strings.TrimSpace(string(output))) > 0
}

// Checks if the required ports specified in the configuration are available.
func verifyPorts() error {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return util.HandleError("Failed to load configuration", err)
	}

	// Check required ports
	requiredPorts := []string{
		cfg.Proxy.HTTPPort,
		cfg.Proxy.HTTPSPort,
		cfg.Proxy.AdminPort,
		cfg.SSH.Port,
	}

	for _, port := range requiredPorts {
		if isPortOpen(port) {
			return util.HandleError("Port is already in use: "+port, nil)
		}
	}
	return nil
}

// isPortOpen checks if a given port is currently in use on localhost
// Returns true if the port is open (in use), false otherwise
func isPortOpen(port string) bool {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// prepareDockerComposeCmd creates a properly configured exec.Cmd for Docker Compose operations
// Includes all necessary environment variables and working directory settings
func prepareDockerComposeCmd(command string, cmdArgs []string, proxyConfigDir string, cfg *config.Config) *exec.Cmd {
	cmd := exec.Command(command, cmdArgs...)
	cmd.Dir = proxyConfigDir
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "COMPOSE_IGNORE_ORPHANS=1")
	cmd.Env = append(cmd.Env, "DOCKER_SOCK="+cfg.Docker.Sock)
	cmd.Env = append(cmd.Env, "DOCKER_PROJECT_NAME="+cfg.Docker.ProjectName)
	cmd.Env = append(cmd.Env, "DOCKER_NETWORK_NAME="+cfg.Docker.NetworkName)
	cmd.Env = append(cmd.Env, "DOCKER_IMAGE_PROXY="+cfg.Proxy.ImageName)
	cmd.Env = append(cmd.Env, "HTTP_PORT="+cfg.Proxy.HTTPPort)
	cmd.Env = append(cmd.Env, "HTTPS_PORT="+cfg.Proxy.HTTPSPort)
	cmd.Env = append(cmd.Env, "ADMIN_PORT="+cfg.Proxy.AdminPort)

	return cmd
}
