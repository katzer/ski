package main

import (
	"fmt"
	"strings"
)

const planetLength int = 21

func formatAndPrint(planets []Planet, opts *Opts) {
	formatter := Formatter{}
	formatter.init()
	var formatted string

	for _, entry := range planets {
		formatted = formatter.format(entry, opts)
		fmt.Print(formatted)

	}
	if opts.Pretty {
		formatter.execute(opts)
	}
}

func makeLoadCommand(command string, opts *Opts) string {
	if opts.Load {
		return fmt.Sprintf(`sh -lc "echo -----APPPLANT-ORBIT----- && %s "`, command)
	}
	return command
}

func cleanProfileLoadedOutput(output string, opts *Opts) string {
	if opts.Load {
		splitOut := strings.Split(output, "-----APPPLANT-ORBIT-----\n")
		return splitOut[len(splitOut)-1]
	}
	return output
}
