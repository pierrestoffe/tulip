package helpers

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/pierrestoffe/tulip/pkg/constants"
    "github.com/pierrestoffe/tulip/pkg/util/log"
)

func GetTulipDir() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", HandleError("error getting home directoryw", err)
    }
    return filepath.Join(homeDir, constants.AppRootDir), nil
}

func GetProxyConfigDir() (string, error) {
    tulipDir, err := GetTulipDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(tulipDir, constants.AppConfigDir, constants.ProxyConfigDir), nil
}

func HandleError(message string, errMessage error) error {
	err := fmt.Errorf(message, errMessage)
	log.PrintError(message + ": " + errMessage.Error())
	return err
}

func CreateDir(dirPath string, perm os.FileMode) error {
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        if err := os.MkdirAll(dirPath, perm); err != nil {
            return err
        }
    }
    return nil
}

func EnsureDir(dirPath string) error {
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        return fmt.Errorf("directory does not exist: %s", dirPath)
    }
    return nil
}
