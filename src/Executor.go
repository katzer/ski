package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Executor This struct ensures the parallel execution of all command executions
type Executor struct {
	planets []Planet
}

func (executor *Executor) execMain(opts *Opts) {
	outputList := make([]StructuredOuput, len(executor.planets))
	var wg sync.WaitGroup

	if opts.debugFlag {
		fmt.Println("prettyflag " + strconv.FormatBool(opts.prettyFlag))
		fmt.Println("scriptflag " + strconv.FormatBool(opts.scriptFlag))
		fmt.Println("command " + opts.command)
		for _, planet := range opts.planets {
			fmt.Printf("planet %s", planet)
		}
	}

	wg.Add(len(executor.planets))
	for i, planet := range executor.planets {
		// to avoid closure over the value planet and the value i. seems odd but it is recommended
		a := i
		planet := planet
		go func() {
			planet.execute(opts)
			outputList[a] = planet.outputStruct
			wg.Done()
		}()
	}

	wg.Wait()
	formatAndPrint(outputList, opts)
}
