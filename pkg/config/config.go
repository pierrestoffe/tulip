// Package config manages the application's configuration settings and provides
// functions to load, validate, and access configuration values
package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/pierrestoffe/tulip/pkg/util"
	"gopkg.in/yaml.v3"
)

// Application constants define paths, versions, and file names used throughout the application
const (
	// App-level constants
	AppName    = "Tulip"  // Name of the application
	AppVersion = "1.0.0"  // Current version of the application
	AppRootDir = ".tulip" // Root directory for all Tulip configuration files

	// Configuration-related constants
	ConfigFile          = "config.yml" // Main configuration file name
	ConfigVersion       = "1.0"        // Configuration file version
	ConfigCertsDir      = "certs"      // Directory for SSL certificates
	ConfigContainersDir = "containers" // Directory for container configurations

	// Proxy-related constants
	ProxyContainerName     = "tulip-proxy"        // Name of the proxy container
	ProxyConfigDir         = "proxy"              // Directory for proxy configuration
	ProxyDockerComposeFile = "docker-compose.yml" // Docker Compose file for proxy
	ProxyTraefikFile       = "traefik.yml"        // Traefik configuration file

	// SSH-related constants
	SSHContainerName     = "tulip-ssh"          // Name of the SSH container
	SSHConfigDir         = "ssh"                // Directory for SSH configuration
	SSHDockerComposeFile = "docker-compose.yml" // Docker Compose file for SSH
	SSHDockerFile        = "Dockerfile"         // Dockerfile for SSH service
)

// Config represents the application configuration
type Config struct {
	Docker DockerConfig `yaml:"docker"`
	Proxy  ProxyConfig  `yaml:"proxy"`
	SSH    SSHConfig    `yaml:"ssh"`
}

// DockerConfig holds Docker-related configuration
type DockerConfig struct {
	Sock        string `yaml:"sock"`
	ProjectName string `yaml:"projectName"`
	NetworkName string `yaml:"networkName"`
}

// ProxyConfig holds proxy-related configuration
type ProxyConfig struct {
	ImageName string `yaml:"imageName"`
	HTTPPort  string `yaml:"httpPort"`
	HTTPSPort string `yaml:"httpsPort"`
	AdminPort string `yaml:"adminPort"`
}

// SSHConfig holds SSH-related configuration
type SSHConfig struct {
	ImageName string `yaml:"imageName"`
	Port      string `yaml:"port"`
}

var (
	// Global configuration instance
	config      *Config
	configMutex sync.RWMutex
)

// DefaultConfig returns a new Config instance with default values.
func DefaultConfig() *Config {
	return &Config{
		Docker: DockerConfig{
			Sock:        "/var/run/docker.sock",
			ProjectName: "tulip",
			NetworkName: "tulip",
		},
		Proxy: ProxyConfig{
			ImageName: "traefik:3.3.4",
			HTTPPort:  "80",
			HTTPSPort: "443",
			AdminPort: "8850",
		},
		SSH: SSHConfig{
			ImageName: "ssh",
			Port:      "8851",
		},
	}
}

// Load reads the configuration file and returns a Config struct
func Load(initialize bool) (*Config, error) {
	var configFileData []byte

	configMutex.Lock() // Only one goroutine can modify at a time
	defer configMutex.Unlock()

	if !initialize && config != nil {
		return config, nil // Return cached config if already loaded
	}

	// Set default configuration
	config = DefaultConfig()

	// Check if config file exists
	tulipDirPath := GetTulipDirPath()
	if _, err := os.Stat(tulipDirPath); !os.IsNotExist(err) {
		tulipDir, err := GetTulipDir()
		if err != nil {
			return nil, err
		}
		configPath := filepath.Join(tulipDir, ConfigFile)

		// Check if config file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			return config, nil
		}

		// Read config file
		configFileData, err = os.ReadFile(configPath)
		if err != nil {
			return nil, util.HandleError("Failed to read configuration file", err)
		}
	}

	// Parse YAML
	if len(configFileData) > 0 {
		if err := yaml.Unmarshal(configFileData, config); err != nil {
			return nil, util.HandleError("Failed to parse configuration file", err)
		}
	}

	// Make sure that key configuration values are valid
	if err := validateConfig(config); err != nil {
		return nil, util.HandleError("Invalid configuration", err)
	}

	return config, nil
}

// Get returns the current configuration
func Get() (*Config, error) {
	configMutex.RLock()
	if config != nil {
		defer configMutex.RUnlock()
		return config, nil
	}
	configMutex.RUnlock()

	return Load(false)
}

// Initialize initializes configuration
func Initialize() (*Config, error) {
	return Load(true)
}

// validateConfig ensures the configuration has valid values
func validateConfig(cfg *Config) error {
	if cfg.Docker.ProjectName == "" {
		return util.HandleError("Docker project name cannot be empty", nil)
	}
	if cfg.Docker.NetworkName == "" {
		return util.HandleError("Docker network name cannot be empty", nil)
	}
	if cfg.Docker.Sock == "" {
		return util.HandleError("Docker socket path cannot be empty", nil)
	}
	if cfg.Proxy.ImageName == "" {
		return util.HandleError("Proxy image name cannot be empty", nil)
	}
	if cfg.SSH.ImageName == "" {
		return util.HandleError("SSH image name cannot be empty", nil)
	}

	return nil
}
