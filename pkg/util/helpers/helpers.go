package helpers

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/print"
)

// Returns the path to the application's root directory
func GetTulipDir() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", HandleError("error getting home directory", err)
    }
    return filepath.Join(homeDir, constants.AppRootDir), nil
}

// Returns the path to the proxy configuration directory
func GetProxyConfigDir() (string, error) {
    tulipDir, err := GetTulipDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(tulipDir, constants.AppConfigDir, constants.ProxyConfigDir), nil
}

// Logs an error and returns a formatted error
func HandleError(message string, err error, additionalContext ...string) error {
    var formattedErr error

    if len(additionalContext) > 0 && additionalContext[0] != "" {
        // Format with additional context (like stderr output)
        formattedErr = fmt.Errorf("%s: %w\n%s", message, err, additionalContext[0])
        print.Error(message + ": " + err.Error() + "\n" + additionalContext[0])
    } else {
        formattedErr = fmt.Errorf("%s: %w", message, err)
        print.Error(message + ": " + err.Error())
    }

    return formattedErr
}

// Creates a directory if it doesn't exist
func CreateDir(dirPath string, perm os.FileMode) error {
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return os.MkdirAll(dirPath, perm)
    }
    return nil
}

// Checks if a directory exists and returns an error if it doesn't
func EnsureDir(dirPath string) error {
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return fmt.Errorf("directory does not exist: %s", dirPath)
    }
    return nil
}
