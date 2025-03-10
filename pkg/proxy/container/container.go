package container

import (
    "os/exec"
    "strings"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/util/log"
)

func isRunning() bool {
    cmd := exec.Command("docker", "ps", "--filter", "name=" + constants.ProxyContainerName, "--format", "{{.Names}}")
    output, err := cmd.Output()
    if err != nil {
        return false
    }

    return len(strings.TrimSpace(string(output))) > 0
}

func Start() error {
    if isRunning() {
        log.PrintWarning("Proxy " + constants.ProxyContainerName + " is running already.")
        log.PrintInfo("You can access the dashboard at " + constants.ProxyUrl)
        return nil
    }

    log.PrintInfo("Starting " + constants.ProxyContainerName + " proxy...")

    tulipDir, err := helpers.GetProxyConfigDir()
    if err != nil {
        return err
    }

    cmd := exec.Command("docker", "compose", "up", "-d")
    cmd.Dir = tulipDir
    if err := cmd.Run(); err != nil {
        helpers.HandleError("error starting " + constants.ProxyContainerName + " proxy", err)
    }
    log.PrintSuccessReplace("Proxy " + constants.ProxyContainerName + " started.")
    log.PrintInfo("You can access the dashboard at " + constants.ProxyUrl)
    return nil
}

func Stop() error {
    if !isRunning() {
        log.PrintWarning("Proxy " + constants.ProxyContainerName + " was stopped already.")
        return nil
    }

    log.PrintInfo("Stopping " + constants.ProxyContainerName + " proxy...")

    tulipDir, err := helpers.GetProxyConfigDir()
    if err != nil {
        return err
    }

    cmd := exec.Command("docker", "compose", "down")
    cmd.Dir = tulipDir
    if err := cmd.Run(); err != nil {
        helpers.HandleError("error stopping " + constants.ProxyContainerName + " proxy", err)
    }
    log.PrintSuccessReplace("Proxy " + constants.ProxyContainerName + " was stopped.")
    return nil
}
