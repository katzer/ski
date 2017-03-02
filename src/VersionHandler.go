package main

import (
	"fmt"
	"os"
)

var version = "undefined"

func printVersion() {
	runtimeOS := getOS()
	progArch := getArch()
	archOS := getOSArch()
	vers := fmt.Sprintf("ski version %s %s %s (%s)", version, progArch, runtimeOS, archOS)
	os.Stdout.WriteString(vers + "\n")
}
