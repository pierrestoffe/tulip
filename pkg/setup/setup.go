// Package setup handles the initialization and verification of Tulip's environment
package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/proxy"
	proxySetup "github.com/pierrestoffe/tulip/pkg/setup/proxy"
	sshSetup "github.com/pierrestoffe/tulip/pkg/setup/ssh"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Contains the template for the config
const configTemplate = `version: {{.Version}}

docker:
    sock: {{.DockerSock}}
    projectName: {{.ProjectName}}
    networkName: {{.NetworkName}}
proxy:
    imageName: {{.ProxyImageName}}
    httpPort: {{.HTTPPort}}
    httpsPort: {{.HTTPSPort}}
    adminPort: {{.AdminPort}}
ssh:
    imageName: {{.SSHImageName}}
    port: {{.SSHPort}}`

// Initializes the Tulip application environment
// It creates necessary directories, extracts configuration files,
// and starts the proxy service. If Tulip is already initialized,
// it prompts the user for confirmation before reinitializing.
func Initialize() error {
	tulipDir, err := config.GetTulipDir()
	if err == nil {
		util.PrintWarning("Tulip is already initialized at " + tulipDir)
		util.PrintWarning("Do you want to reinitialize? This will overwrite existing configuration. (y/N)")

		// Get user confirmation
		var confirm string
		if _, err := fmt.Scanln(&confirm); err != nil {
			return util.HandleError("Failed to read user input", err)
		}
		util.PrintEmpty()

		if confirm != "y" && confirm != "Y" {
			if err := proxy.Start(); err != nil {
				return err
			}
			return nil
		}
	}

	util.PrintInfo("Initializing Tulip..")

	if err := addConfigFiles(tulipDir); err != nil {
		return err
	}

	util.PrintSuccess("Tulip initialized successfully!")
	util.PrintEmpty()

	// Start or restart the proxy service
	if err := proxy.Restart(); err != nil {
		return err
	}
	return nil
}

// Ensure checks if all required directories and files exist
func Ensure() error {
	// Check required directories
	requiredDirs := []string{
		config.GetTulipDirPath(),
		config.GetContainersConfigDirPath(),
		config.GetCertsConfigDirPath(),
		config.GetProxyConfigDirPath(),
		config.GetSSHConfigDirPath(),
	}

	for _, dir := range requiredDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			util.HandleError("Directory missing: "+dir, nil)
			util.PrintWarning("Please run 'tulip init' to repair")
			util.PrintWarning("Note that the files in ~/.tulip may be overwritten in the process")
			return err
		} else if err != nil {
			return util.HandleError("Error accessing directory: "+dir, nil)
		}
	}

	// Check required files
	requiredFiles := []string{
		filepath.Join(config.GetTulipDirPath(), config.ConfigFile),
		filepath.Join(config.GetProxyConfigDirPath(), config.ProxyDockerComposeFile),
		filepath.Join(config.GetProxyConfigDirPath(), config.ProxyTraefikFile),
		filepath.Join(config.GetSSHConfigDirPath(), config.SSHDockerComposeFile),
		filepath.Join(config.GetSSHConfigDirPath(), config.SSHDockerFile),
	}

	for _, file := range requiredFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			util.HandleError("Required file missing: "+file, nil)
			util.PrintWarning("Please run 'tulip init' to repair")
			util.PrintWarning("Note that the files in ~/.tulip may be overwritten in the process")
			return err
		} else if err != nil {
			return util.HandleError("Error accessing file: "+file, nil)
		}
	}

	return nil
}

// addConfigFiles creates all necessary configuration files in the specified directory
// Returns an error if any file creation or directory setup fails
func addConfigFiles(tulipHomePath string) error {
	// Get configuration
	cfg, err := config.Initialize()
	if err != nil {
		return util.HandleError("Failed to load configuration", err)
	}

	// Prepare template data
	templateData := map[string]string{
		"Version":        config.ConfigVersion,
		"DockerSock":     cfg.Docker.Sock,
		"ProjectName":    cfg.Docker.ProjectName,
		"NetworkName":    cfg.Docker.NetworkName,
		"ProxyImageName": cfg.Proxy.ImageName,
		"HTTPPort":       cfg.Proxy.HTTPPort,
		"HTTPSPort":      cfg.Proxy.HTTPSPort,
		"AdminPort":      cfg.Proxy.AdminPort,
		"SSHImageName":   cfg.SSH.ImageName,
		"SSHPort":        cfg.SSH.Port,
	}

	// Create directories
	for _, dir := range []string{
		filepath.Join(tulipHomePath, config.ConfigContainersDir),
		filepath.Join(tulipHomePath, config.ConfigCertsDir),
	} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return util.HandleError("Error creating directory"+dir, err)
		}
	}

	// Create config.yml
	configFilePath := filepath.Join(tulipHomePath, config.ConfigFile)
	if err := util.CreateFileFromTemplate(configFilePath, configTemplate, templateData); err != nil {
		return err
	}

	// Add setup files
	if err := proxySetup.Initialize(); err != nil {
		return err
	}
	if err := sshSetup.Initialize(); err != nil {
		return err
	}

	return nil
}
