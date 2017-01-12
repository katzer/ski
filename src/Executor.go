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
		//fmt.Println(args)
		fmt.Println("prettyflag " + strconv.FormatBool(opts.prettyFlag))
		fmt.Println("scriptflag " + strconv.FormatBool(opts.scriptFlag))
		fmt.Println("scriptpath " + opts.scriptPath)
		fmt.Println("command " + opts.command)
		for _, planet := range opts.planets {
			fmt.Printf("planet %s", planet)
		}
	}

	wg.Add(len(executor.planets))
	for i, planet := range executor.planets {
		// to avoid closure over the value planet. seems odd but it is recommended
		a := i
		planet := planet
		go func() {
			planet.execute(opts)
			outputList[a] = planet.outputStruct
			wg.Done()
		}()
	}

	/*for i, planet := range opts.planets {
		if opts.typeFlag {
			fmt.Println("The type of " + planet + " is " + getType(planet))
		}

		switch getType(planet) {
		case "server":
			connDet := getConnDet(planet)
			opts.currentDet = connDet
			outputList[i].planet = planet
			if opts.scriptFlag {
				go exec.upAndExecSSHSync(connDet, &wg, &outputList[i], opts)
			} else {
				go exec.execSSHSync(connDet, opts.command, &wg, &outputList[i], opts)
			}
		case "db":
			dbDet := getConnDet(planet)
			opts.currentDBDet = dbDet
			outputList[i].planet = planet
			if opts.scriptFlag {
				go exec.upAndExecDBSync(dbDet, &wg, &outputList[i], opts)
			} else {
				go exec.execDBSync(dbDet, &wg, &outputList[i], opts)
			}
		case "web":
			fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
			os.Exit(1)
		default:
			fmt.Println("default did done")
			wg.Done()
		}
	}*/
	wg.Wait()
	formatAndPrint(outputList, opts)
}
