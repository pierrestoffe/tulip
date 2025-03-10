package proxy

import (
    proxyContainer "github.com/pierrestoffe/tulip/pkg/proxy/container"
    proxyNetwork "github.com/pierrestoffe/tulip/pkg/proxy/network"
)

// Initializes both the proxy network and container
// Returns an error if either component fails to start
func Start() error {
    if err := proxyNetwork.Start(); err != nil {
        return err
    }
    return proxyContainer.Start()
}

// Terminates both the proxy container and network
// Returns an error if either component fails to stop
func Stop() error {
    if err := proxyContainer.Stop(); err != nil {
        return err
    }
    return proxyNetwork.Stop()
}

// Stops and then starts the proxy service
// Returns an error if either stop or start operations fail
func Restart() error {
    if err := Stop(); err != nil {
        return err
    }
    return Start()
}
