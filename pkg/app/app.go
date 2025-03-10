package app

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/pierrestoffe/tulip/pkg/util/log"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/proxy"
    proxyFiles "github.com/pierrestoffe/tulip/pkg/proxy/files"
)

func Initialize() {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        helpers.HandleError("error getting home directory: ", err)
        return
    }

    tulipHomePath := filepath.Join(homeDir, constants.AppRootDir)

    if err := helpers.EnsureDir(tulipHomePath); err == nil {
        log.PrintWarning("Tulip is already initialized at " + tulipHomePath)
        log.PrintWarning("Do you want to reinitialize? This will overwrite existing configuration. (y/N)")

        var confirm string
        fmt.Scanln(&confirm)
        if confirm != "y" && confirm != "Y" {
            log.PrintEmpty()
            proxy.Start()
            return
        }
    }

    log.PrintEmpty()
    log.PrintInfo("Initializing Tulip...")

    if err := helpers.CreateDir(filepath.Join(tulipHomePath, constants.AppConfigDir), 0755); err != nil {
        helpers.HandleError("Error creating directory structure:", err)
        return
    }

    proxyFiles.ExtractConfigFiles()

    log.PrintSuccess("Tulip initialized successfully!")
    log.PrintEmpty()

    proxy.Start()
}
