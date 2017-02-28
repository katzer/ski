package main

import (
	log "github.com/Sirupsen/logrus"
	"strconv"
	"sync"
)

// Executor This struct ensures the parallel execution of all command executions
type Executor struct {
	planets []Planet
}

func (executor *Executor) execMain(opts *Opts) {
	log.Debugf("Function: execMain")
	outputList := make([]StructuredOuput, len(executor.planets))
	var wg sync.WaitGroup

	log.Debugln("pretty " + strconv.FormatBool(opts.pretty))
	log.Debugln("script " + opts.scriptName)
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
