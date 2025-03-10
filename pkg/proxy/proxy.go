package proxy

import (
    proxyContainer "github.com/pierrestoffe/tulip/pkg/proxy/container"
    proxyNetwork "github.com/pierrestoffe/tulip/pkg/proxy/network"
)

func Start() error {
    if err := proxyNetwork.Start(); err != nil {
        return err
    }
    return proxyContainer.Start()
}

func Stop() error {
    if err := proxyContainer.Stop(); err != nil {
        return err
    }
    return proxyNetwork.Stop()
}
