// Package config provides functionality for managing configuration directories
// and file paths used by the Tulip application
package config

import (
	"os"
	"path/filepath"

	"github.com/pierrestoffe/tulip/pkg/util"
)

// GetUserHomeDir retrieves the user's home directory path
// Returns an error if the home directory cannot be determined
func GetUserHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", util.HandleError("Error getting home directory", err)
	}
	return homeDir, nil
}

// GetTulipDirPath constructs the full path to Tulip's configuration directory
func GetTulipDirPath() string {
	homeDir, err := GetUserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(homeDir, AppRootDir)
}

// GetTulipDir verifies and returns the path to Tulip's configuration directory
// Returns an error if the directory doesn't exist
func GetTulipDir() (string, error) {
	tulipDirPath := GetTulipDirPath()

	// Check if directory exists
	if _, err := os.Stat(tulipDirPath); os.IsNotExist(err) {
		return "", util.HandleError("Tulip home directory does not exist", err)
	}

	return tulipDirPath, nil
}

// GetCertsConfigDirPath constructs the full path to the certificates directory
func GetCertsConfigDirPath() string {
	return filepath.Join(GetTulipDirPath(), ConfigCertsDir)
}

// GetCertsConfigDir verifies and returns the path to the certificates directory
// Returns an error if the directory doesn't exist
func GetCertsConfigDir() (string, error) {
	certsConfigDirPath := GetCertsConfigDirPath()

	// Check if directory exists
	if _, err := os.Stat(certsConfigDirPath); os.IsNotExist(err) {
		return "", util.HandleError("Certs config directory does not exist", err)
	}

	return certsConfigDirPath, nil
}

// GetContainersConfigDirPath constructs the full path to the containers configuration directory
func GetContainersConfigDirPath() string {
	return filepath.Join(GetTulipDirPath(), ConfigContainersDir)
}

// GetContainersConfigDir verifies and returns the path to the containers configuration directory
// Returns an error if the directory doesn't exist
func GetContainersConfigDir() (string, error) {
	containersConfigDirPath := GetContainersConfigDirPath()

	// Check if directory exists
	if _, err := os.Stat(containersConfigDirPath); os.IsNotExist(err) {
		return "", util.HandleError("Containers config directory does not exist", err)
	}

	return containersConfigDirPath, nil
}

// GetProxyConfigDirPath constructs the full path to the proxy configuration directory
func GetProxyConfigDirPath() string {
	return filepath.Join(GetContainersConfigDirPath(), ProxyConfigDir)
}

// GetProxyConfigDir verifies and returns the path to the proxy configuration directory
// Returns an error if the directory doesn't exist
func GetProxyConfigDir() (string, error) {
	proxyConfigDirPath := GetProxyConfigDirPath()

	// Check if directory exists
	if _, err := os.Stat(proxyConfigDirPath); os.IsNotExist(err) {
		return "", util.HandleError("Proxy config directory does not exist", err)
	}

	return proxyConfigDirPath, nil
}

// GetSSHConfigDirPath constructs the full path to the SSH configuration directory
func GetSSHConfigDirPath() string {
	return filepath.Join(GetContainersConfigDirPath(), SSHConfigDir)
}

// GetSSHConfigDir verifies and returns the path to the SSH configuration directory
// Returns an error if the directory doesn't exist
func GetSSHConfigDir() (string, error) {
	sshConfigDirPath := GetSSHConfigDirPath()

	// Check if directory exists
	if _, err := os.Stat(sshConfigDirPath); os.IsNotExist(err) {
		return "", util.HandleError("SSH config directory does not exist", err)
	}

	return sshConfigDirPath, nil
}
