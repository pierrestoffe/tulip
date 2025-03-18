// Package proxy provides functionality for configuring and managing the Traefik proxy service
package proxy

import (
	"os"
	"path/filepath"

	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Contains the template for the Docker Compose file
const dockerComposeTemplate = `name: ${DOCKER_PROJECT_NAME}

services:
  proxy:
    image: ${DOCKER_IMAGE_PROXY}
    container_name: tulip-proxy
    restart: unless-stopped
    networks:
      - tulip-default
    ports:
      - "${HTTP_PORT}:80"
      - "${HTTPS_PORT}:443"
      - "${ADMIN_PORT}:8080"
    volumes:
      - ./traefik.yml:/etc/traefik/traefik.yml:ro
      - ${CONFIG_ROOT:-./../..}/certs/:/etc/traefik/certs/:ro
      - ${DOCKER_SOCK:-/var/run/docker.sock}:/var/run/docker.sock:ro

networks:
  tulip-default:
    name: ${DOCKER_NETWORK_NAME}
    external: true`

// Contains the template for the Traefik configuration file
const traefikTemplate = `api:
  dashboard: true
  insecure: true

entryPoints:
  web:
    address: ":80"
  websecure:
    address: ":443"

providers:
  docker:
    endpoint: "unix:///var/run/docker.sock"
    exposedByDefault: false
    network: traefik
  file:
    directory: "/etc/traefik/certs/"
    watch: true

log:
  level: "DEBUG"

accessLog:
  filePath: "/var/log/traefik/access.log"
  format: json`

// Initialize creates the necessary proxy configuration files
// It sets up docker-compose.yml and traefik.yml with the proper configuration
// Returns an error if any file creation fails
func Initialize() error {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		return util.HandleError("Failed to load configuration", err)
	}

	// Construct the path to Tulip's proxy directory
	proxyConfigDirPath := config.GetProxyConfigDirPath()

	// Create proxy directory if it doesn't exist
	if err := os.MkdirAll(proxyConfigDirPath, 0755); err != nil {
		return util.HandleError("Failed to create proxy directory", err)
	}

	// Prepare template data
	// TODO: remove?
	templateData := map[string]string{
		"HTTPPort":  cfg.Proxy.HTTPPort,
		"HTTPSPort": cfg.Proxy.HTTPSPort,
		"AdminPort": cfg.Proxy.AdminPort,
	}

	// Create docker-compose.yml
	dockerComposePath := filepath.Join(proxyConfigDirPath, config.ProxyDockerComposeFile)
	if err := util.CreateFileFromTemplate(dockerComposePath, dockerComposeTemplate, templateData); err != nil {
		return err
	}

	// Create traefik.yml
	traefikPath := filepath.Join(proxyConfigDirPath, config.ProxyTraefikFile)
	if err := util.CreateFileFromTemplate(traefikPath, traefikTemplate, templateData); err != nil {
		return err
	}
	return nil
}
