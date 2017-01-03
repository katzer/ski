package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Executor struct {
}

func (exec *Executor) execMain(opts *Opts) {
	outputList := make([]StructuredOuput, len(opts.planets))
	var wg sync.WaitGroup

	if opts.debugFlag {
		//println(args)
		println("prettyflag " + strconv.FormatBool(opts.prettyFlag))
		println("scriptflag " + strconv.FormatBool(opts.scriptFlag))
		println("scriptpath " + opts.scriptPath)
		println("command " + opts.command)
		for _, planet := range opts.planets {
			println("planet " + planet)
		}
	}

	wg.Add(len(opts.planets))
	for i, planet := range opts.planets {
		if opts.typeFlag {
			println("The type of " + planet + " is " + getType(planet))
		}

		switch getType(planet) {
		case "server":
			connDet := opts.getConnDet(planet)
			opts.currentDet = connDet
			outputList[i].planet = planet
			if opts.scriptFlag {
				go exec.upAndExecSSHSync(connDet, &wg, &outputList[i], opts)
			} else {
				go exec.execSSHSync(connDet, opts.command, &wg, &outputList[i], opts)
			}
		case "db":
			dbDet := opts.getConnDet(planet)
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
			println("default did done")
			wg.Done()
		}
	}
	wg.Wait()
	formatAndPrint(outputList, opts)
}

func (exec *Executor) execSSHSync(connDet string, command string, wg *sync.WaitGroup, outputStruct *StructuredOuput, opts *Opts) {
	execSSHCommand(connDet, opts.command, outputStruct, opts)
	wg.Done()
}

func (exec *Executor) upAndExecSSHSync(connDet string, wg *sync.WaitGroup, outputStruct *StructuredOuput, opts *Opts) {
	upAndExecSSHScript(connDet, outputStruct, opts)
	wg.Done()
}

func (exec *Executor) execDBSync(dbDet string, wg *sync.WaitGroup, outputStruct *StructuredOuput, opts *Opts) {
	execDBCommand(dbDet, outputStruct, opts)
	wg.Done()
}

func (exec *Executor) upAndExecDBSync(dbDet string, wg *sync.WaitGroup, outputStruct *StructuredOuput, opts *Opts) {
	upAndExecDBScript(dbDet, outputStruct, opts)
	wg.Done()
}
