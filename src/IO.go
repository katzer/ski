package main

import (
	"fmt"
	"os"
	"strings"
)

func formatAndPrint(planets []Planet, opts *Opts) {
	formatter := Formatter{}
	failed := false
	formatter.init()
	var formatted string
	for _, entry := range planets {
		if entry.outputStruct.errored {
			failed = true
		}
		formatted = formatter.format(entry, opts)
		fmt.Print(formatted)

	}
	if opts.Pretty {
		formatter.execute(opts)
	}
	if failed {
		os.Exit(1)
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
