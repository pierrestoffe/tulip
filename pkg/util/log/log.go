package log

import (
    "fmt"
)

func PrintInfo(message string) {
    fmt.Printf("%s\n", message)  // White
}

func PrintSuccess(message string) {
    fmt.Printf("\033[32m%s\033[0m\n", message)  // Green
}

func PrintWarning(message string) {
    fmt.Printf("\033[33m%s\033[0m\n", message)  // Yellow
}

func PrintError(message string) {
    fmt.Printf("\033[31m%s\033[0m\n", message)  // Red
}

func PrintDebug(message string) {
    fmt.Printf("\033[34m%s\033[0m\n", message)  // Blue
}
