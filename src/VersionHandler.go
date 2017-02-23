package main

import (
	"fmt"
	"os"
)

var version = "undefined"

/**
*	Prints the current Version of the ski application
 */
func printVersion() {
	runtimeOS := getOS()
	progArch := getArch()
	archOS := getOSArch()
	vers := fmt.Sprintf("ski version %s %s %s (%s)", version, progArch, runtimeOS, archOS)
	os.Stdout.WriteString(vers + "\n")
}
