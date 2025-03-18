// Package project handles project-level operations in Tulip
package project

import "github.com/pierrestoffe/tulip/pkg/util"

// Start begins the execution of a Tulip project in the current directory
func Start() error {
	util.PrintDebug("hello")
	return nil
}
