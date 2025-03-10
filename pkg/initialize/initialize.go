package initialize

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/pierrestoffe/tulip/pkg/util/log"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/constants"
    proxyFiles "github.com/pierrestoffe/tulip/pkg/proxy/files"
)

const (
    version = "1.0.0"
    tulipHome = ".tulip"
    tulipTemplates = ".tulip/templates"
    traefikDir = ".tulip/traefik"
    configFile = ".tulip/config.yml"
)

func InitTulip() {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        helpers.HandleError("error getting home directory: ", err)
        return
    }

    tulipHomePath := filepath.Join(homeDir, tulipHome)

    if err := helpers.VerifyDir(tulipHomePath); err == nil {
        log.PrintWarning("Tulip is already initialized at " + tulipHomePath)
        log.PrintWarning("Do you want to reinitialize? This will overwrite existing configuration. (y/N)")

        var confirm string
        fmt.Scanln(&confirm)
        if confirm != "y" && confirm != "Y" {
            log.PrintSuccess("You can use 'tulip start' to start Tulip's proxy.")
            return
        }
    }

    fmt.Println("\nInitializing Tulip...")

    if err := helpers.CreateDir(filepath.Join(tulipHomePath, constants.AppConfigDir), 0755); err != nil {
        fmt.Println("Error creating directory structure:", err)
        return
    }

    proxyFiles.ExtractConfigFiles()

    log.PrintSuccess("Tulip initialized successfully!")
    log.PrintSuccess("You can now use 'tulip start' to start Tulip's proxy.")
}
