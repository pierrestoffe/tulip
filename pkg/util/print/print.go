package print

import (
    "fmt"
)

const (
    colorReset  = "\033[0m"  // ANSI code to reset text color
    colorRed    = "\033[31m" // ANSI code for red text
    colorGreen  = "\033[32m" // ANSI code for green text
    colorYellow = "\033[33m" // ANSI code for yellow text
    colorBlue   = "\033[34m" // ANSI code for blue text
    colorWhite  = ""         // Default terminal color
)

// Prints a message in the default terminal color
func Info(message string) {
    print(message, colorWhite)
}

// Prints a message in green color to indicate success
func Success(message string) {
    print(message, colorGreen)
}

// Prints a message in yellow color to indicate a warning
func Warning(message string) {
    print(message, colorYellow)
}

// Prints a message in red color to indicate an error
func Error(message string) {
    print(message, colorRed)
}

// Prints a message in blue color for debug information
func Debug(message string) {
    print(message, colorBlue)
}

// Clears the previous line and prints a message in the default terminal color
func InfoReplace(message string) {
    printReplace(message, colorWhite)
}

// Clears the previous line and prints a message in green color
func SuccessReplace(message string) {
    printReplace(message, colorGreen)
}

// Clears the previous line and prints a message in yellow color
func WarningReplace(message string) {
    printReplace(message, colorYellow)
}

// Clears the previous line and prints a message in red color
func ErrorReplace(message string) {
    printReplace(message, colorRed)
}

// Clears the previous line and prints a message in blue color
func DebugReplace(message string) {
    printReplace(message, colorBlue)
}

// Prints an empty line
func Empty() {
    Info("")
}

// Outputs a message with the specified color to stdout
func print(message string, color string) {
    if color == "" {
        fmt.Printf("%s\n", message)
    } else {
        fmt.Printf("%s%s%s\n", color, message, colorReset)
    }
}

// Clears the previous line and prints a message with the specified color
func printReplace(message string, color string) {
    clearLine()
    print(message, color)
}

// Moves the cursor up one line and clears it
func clearLine() {
    fmt.Print("\033[1A\033[K") // Move up one line and clear it
}
