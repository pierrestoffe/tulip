package app

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/pierrestoffe/tulip/pkg/util/print"
    "github.com/pierrestoffe/tulip/pkg/util/helpers"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/proxy"
    proxyFiles "github.com/pierrestoffe/tulip/pkg/proxy/files"
)

// Initialize sets up the Tulip application environment
// It creates necessary directories, extracts configuration files,
// and starts the proxy service. If Tulip is already initialized,
// it prompts the user for confirmation before reinitializing.
func Initialize() {
    // Get the user's home directory
    homeDir, err := os.UserHomeDir()
    if err != nil {
        helpers.HandleError("error getting home directory: ", err)
        return
    }

    // Construct the path to Tulip's home directory
    tulipHomePath := filepath.Join(homeDir, constants.AppRootDir)

    // Check if Tulip is already initialized
    if err := helpers.EnsureDir(tulipHomePath); err == nil {
        print.Warning("Tulip is already initialized at " + tulipHomePath)
        print.Warning("Do you want to reinitialize? This will overwrite existing configuration. (y/N)")

        // Get user confirmation
        var confirm string
        fmt.Scanln(&confirm)
        if confirm != "y" && confirm != "Y" {
            print.Empty()
            proxy.Start()
            return
        }
    }

    print.Empty()
    print.Info("Initializing Tulip...")

    // Create the configuration directory structure
    if err := helpers.CreateDir(filepath.Join(tulipHomePath, constants.AppConfigDir), 0755); err != nil {
        helpers.HandleError("Error creating directory structure:", err)
        return
    }

    // Add proxy configuration files
    proxyFiles.AddConfigFiles()

    print.Success("Tulip initialized successfully!")
    print.Empty()

    // Start the proxy service
    proxy.Start()
}
