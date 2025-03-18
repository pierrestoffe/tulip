// Package util provides utility functions for output formatting, error handling,
// and file operations used throughout the Tulip application
package util

import (
	"fmt"
	"strings"
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
func PrintInfo(message string) {
	print(message, colorWhite)
}

// Prints a message in green color to indicate success
func PrintSuccess(message string) {
	print(message, colorGreen)
}

// Prints a message in yellow color to indicate a warning
func PrintWarning(message string) {
	print(message, colorYellow)
}

// Prints a message in red color to indicate an error
func PrintError(message string) {
	print(message, colorRed)
}

// Prints a message in red color to indicate an error and returns the error
func PrintErrorE(err error) error {
	return fmt.Errorf("%s%v%s", colorRed, err, colorReset)
}

// Prints a message in blue color for debug information
func PrintDebug(message string) {
	print(message, colorBlue)
}

// Clears the previous line and prints a message in the default terminal color
func PrintInfoReplace(message string) {
	printReplace(message, colorWhite)
}

// Clears the previous line and prints a message in green color
func PrintSuccessReplace(message string) {
	printReplace(message, colorGreen)
}

// Clears the previous line and prints a message in yellow color
func PrintWarningReplace(message string) {
	printReplace(message, colorYellow)
}

// Clears the previous line and prints a message in red color
func PrintErrorReplace(message string) {
	printReplace(message, colorRed)
}

// Clears the previous line and prints a message in blue color
func PrintDebugReplace(message string) {
	printReplace(message, colorBlue)
}

// Prints an empty line
func PrintEmpty() {
	PrintInfo("")
}

// HandleError formats and prints an error message with optional additional context
// Parameters:
//   - message: The main error message to display
//   - err: The underlying error (may be nil)
//   - additionalContext: Optional extra information to append to the error
//
// Returns a formatted error that includes all provided information
func HandleError(message string, err error, additionalContext ...string) error {
	var formattedErr error
	var displayMessage string

	// Create a lowercase version of the message
	messageLowercase := strings.ToLower(message)

	// Check if we have additional context
	hasContext := len(additionalContext) > 0 && additionalContext[0] != ""

	// Handle different combinations of err and additionalContext
	if err == nil {
		if hasContext {
			// Case: No error but with additional context
			displayMessage = fmt.Sprintf("%s\n%s", message, additionalContext[0])
			formattedErr = fmt.Errorf("%s\n%s", messageLowercase, additionalContext[0])
		} else {
			// Case: No error and no additional context
			displayMessage = message
			formattedErr = fmt.Errorf("%s", messageLowercase)
		}
	} else {
		if hasContext {
			// Case: Error with additional context
			displayMessage = fmt.Sprintf("%s: %v\n%s", message, err, additionalContext[0])
			formattedErr = fmt.Errorf("%s: %v\n%s", messageLowercase, err, additionalContext[0])
		} else {
			// Case: Error without additional context
			displayMessage = fmt.Sprintf("%s: %v", message, err)
			formattedErr = fmt.Errorf("%s: %v", messageLowercase, err)
		}
	}

	// Print the error message
	PrintError(displayMessage)

	return formattedErr
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
