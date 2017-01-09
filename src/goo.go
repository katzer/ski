package main

import (
	"os"
)

// StructuredOuput ...
type StructuredOuput struct {
	planet       string
	output       string
	maxOutLength int
}

/**
*	Main function
 */
func main() {
	args := os.Args
	opts := Opts{}
	exec := Executor{}

	opts.procArgs(args)
	if opts.helpFlag {
		printHelp()
	}
	if opts.versionFlag {
		printVersion()
	}
	exec.execMain(&opts)

}
