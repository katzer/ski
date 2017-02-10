package main

import (
	"strconv"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// Executor This struct ensures the parallel execution of all command executions
type Executor struct {
	planets []Planet
}

func (executor *Executor) execMain(opts *Opts) {
	outputList := make([]StructuredOuput, len(executor.planets))
	var wg sync.WaitGroup

	log.Debugln("prettyflag " + strconv.FormatBool(opts.prettyFlag))
	log.Debugln("scriptflag " + strconv.FormatBool(opts.scriptFlag))
	log.Debugln("command " + opts.command)
	for _, planet := range opts.planets {
		log.Debugf("planet %s", planet)
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
