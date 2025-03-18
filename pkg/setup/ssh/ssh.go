// Package ssh provides functionality for configuring and managing the SSH tunnel service
package ssh

import (
	"os"
	"path/filepath"

	"github.com/pierrestoffe/tulip/pkg/config"
	"github.com/pierrestoffe/tulip/pkg/util"
)

// Contains the template for the Docker Compose file
const dockerComposeTemplate = `name: ${DOCKER_PROJECT_NAME}

services:
  ssh-tunnel:
    build: ./
    container_name: tulip-ssh-tunnel
    restart: unless-stopped
    networks:
      - tulip-default
    ports:
      - "${SSH_PORT}:22"

networks:
  tulip-default:
    name: ${DOCKER_NETWORK_NAME}
    external: true`

// Contains the template for the Dockerfile
const dockerFileTemplate = `FROM alpine:latest

# Install OpenSSH server and MariaDB client
RUN apk add --no-cache openssh mariadb-client \
    && rm -rf /var/cache/apk/*

# Create required directories
RUN mkdir -p /var/run/sshd

# Configure SSH with more permissive settings for clients like TablePlus
RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'ChallengeResponseAuthentication no' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'UsePAM yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'PermitTunnel yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'GatewayPorts yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'AllowTcpForwarding yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'ClientAliveInterval 30' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'ClientAliveCountMax 3' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'TCPKeepAlive yes' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'LogLevel DEBUG3' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'PermitOpen any' >> /etc/ssh/sshd_config.d/custom.conf \
    && echo 'AllowStreamLocalForwarding yes' >> /etc/ssh/sshd_config.d/custom.conf

# Create a tunnel user with a simple password
RUN adduser -D -s /bin/sh tulip \
    && echo "tulip:tulip" | chpasswd

# Generate host keys
RUN ssh-keygen -A

# Expose SSH port
EXPOSE 22

# Start SSH daemon with debugging
CMD ["/usr/sbin/sshd", "-D", "-e"]`

// Initialize creates the necessary SSH configuration files
// It sets up docker-compose.yml and Dockerfile with the proper configuration
// Returns an error if any file creation fails
func Initialize() error {
	// Get configuration
	cfg, err := config.Get()
	if err != nil {
		util.HandleError("Failed to load configuration", err)
	}

	// Construct the path to Tulip's ssh directory
	sshConfigDirPath := config.GetSSHConfigDirPath()

	// Create ssh directory if it doesn't exist
	if err := os.MkdirAll(sshConfigDirPath, 0755); err != nil {
		return util.HandleError("Failed to create ssh directory", err)
	}

	// Prepare template data
	templateData := map[string]string{
		"SSHPort": cfg.SSH.Port,
	}

	// Create docker-compose.yml
	dockerComposePath := filepath.Join(sshConfigDirPath, config.SSHDockerComposeFile)
	if err := util.CreateFileFromTemplate(dockerComposePath, dockerComposeTemplate, templateData); err != nil {
		return err
	}

	// Create Dockerfile
	dockerFilePath := filepath.Join(sshConfigDirPath, config.SSHDockerFile)
	if err := util.CreateFileFromTemplate(dockerFilePath, dockerFileTemplate, templateData); err != nil {
		return err
	}
	return nil
}
