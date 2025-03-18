// Package util provides utility functions used throughout the Tulip application
package util

import (
	"os"
	"path/filepath"
	"text/template"
)

// CreateFileFromTemplate generates a new file from a template with the provided data
// Parameters:
//   - destPath: target path where the file will be created
//   - templateContent: the template string to process
//   - data: key-value pairs to use in template processing
func CreateFileFromTemplate(destPath string, templateContent string, data map[string]string) error {
	// Use empty map if data is nil
	templateData := data
	if templateData == nil {
		templateData = make(map[string]string)
	}

	// Process file as template
	tmpl, err := template.New(filepath.Base(destPath)).Parse(templateContent)
	if err != nil {
		return HandleError("Failed to parse template for"+destPath, err)
	}

	// Create destination file
	f, err := os.Create(destPath)
	if err != nil {
		return HandleError("Failed to create file"+destPath, err)
	}
	defer f.Close()

	// Add content to destination file
	err = tmpl.Execute(f, templateData)
	if err != nil {
		return HandleError("Failed to add content to destination file"+destPath, err)
	}

	PrintInfo("Created " + destPath)
	return nil
}
