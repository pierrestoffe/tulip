package container

import (
    "os/exec"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/util/log"
)

func isRunning() bool {
    cmd := exec.Command("docker", "network", "inspect", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        return false
    }
    return true
}

func Start() error {
    if isRunning() {
        log.PrintWarning("Proxy network " + constants.ProxyNetworkName + " is running already.")
        return nil
    }

    log.PrintInfo("Starting " + constants.ProxyNetworkName + " proxy network...")
    cmd := exec.Command("docker", "network", "create", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        helpers.HandleError("error starting " + constants.ProxyNetworkName + " proxy network", err)
    }
    log.PrintSuccessReplace("Proxy network " + constants.ProxyNetworkName + " started.")
    return nil
}

func Stop() error {
    if !isRunning() {
        log.PrintWarning("Proxy network " + constants.ProxyNetworkName + " was stopped already.")
        return nil
    }

    log.PrintInfo("Stopping " + constants.ProxyNetworkName + " network...")

    cmd := exec.Command("docker", "network", "remove", constants.ProxyNetworkName)
    if err := cmd.Run(); err != nil {
        helpers.HandleError("error stopping " + constants.ProxyNetworkName + " proxy network", err)
    }
    log.PrintSuccessReplace("Proxy network " + constants.ProxyNetworkName + " was stopped.")
    return nil
}
