package files

import (
    "os"
    "path/filepath"
    "text/template"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/util/print"
)

// Contains the template for the Docker Compose configuration
const dockerComposeTemplate = `name: tulip

services:
  proxy:
    image: traefik:v2.10
    container_name: tulip_proxy
    ports:
      - "80:80"
      - "443:443"
      - "8855:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./traefik.yml:/etc/traefik/traefik.yml
    networks:
      - tulip
    restart: unless-stopped

networks:
  tulip:
    external: true`

// Contains the template for the Traefik configuration
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

log:
  level: "INFO"`

// Creates the necessary configuration files for the proxy
// It generates both Docker Compose and Traefik configuration files
func AddConfigFiles() error {
    // Get Tulip's home directory
    tulipDir, err := helpers.GetTulipDir()
    if err != nil {
        return helpers.HandleError("failed to get tulip directory: ", err)
    }

    // Construct the path to Tulip's proxy directory
    proxyConfigDir := filepath.Join(tulipDir, constants.AppConfigDir, constants.ProxyConfigDir)

    // Create proxy directory if it doesn't exist
    if err := os.MkdirAll(proxyConfigDir, 0755); err != nil {
        return helpers.HandleError("failed to create proxy directory: ", err)
    }

    // Create docker-compose.yml
    dockerComposePath := filepath.Join(proxyConfigDir, constants.ProxyDockerFile)
    if err := createFileFromTemplate(dockerComposePath, dockerComposeTemplate); err != nil {
        return err
    }

    // Create traefik.yml
    traefikPath := filepath.Join(proxyConfigDir, constants.ProxyTraefikFile)
    if err := createFileFromTemplate(traefikPath, traefikTemplate); err != nil {
        return err
    }
    return nil
}

// Generates a file at the specified path using the provided template
// It parses the template, creates the file, and writes the processed template content
func createFileFromTemplate(destPath string, templateContent string) error {
    // Process file as template
    tmpl, err := template.New(filepath.Base(destPath)).Parse(templateContent)
    if err != nil {
        return helpers.HandleError("failed to parse template for " + destPath, err)
    }

    // Create destination file
    f, err := os.Create(destPath)
    if err != nil {
        return helpers.HandleError("failed to create file " + destPath, err)
    }
    defer f.Close()

    // Add content to destination file
    err = tmpl.Execute(f, map[string]string{})
    if err != nil {
        return helpers.HandleError("failed to add content to destination file " + destPath, err)
    }

    print.Info("Created " + destPath)
    return nil
}
