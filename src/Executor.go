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
	var wg sync.WaitGroup

	log.Debugln("prettyflag " + strconv.FormatBool(opts.prettyFlag))
	log.Debugln("script " + opts.scriptName)
	log.Debugln("command " + opts.command)
	for _, planet := range opts.planets {
		log.Debugf("planet %s", planet)
	}

	wg.Add(len(executor.planets))

	for _, planet := range executor.planets {
		// to avoid closure over the value planet and the value i. seems odd but it is recommended
		planet := planet
		go func() {
			planet.execute(opts)
			wg.Done()
		}()
	}
	wg.Wait()
	formatAndPrint(executor.planets, opts)
}
