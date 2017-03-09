package main

import (
	"strconv"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// Executor This struct ensures the parallel execution of all command executions
type Executor struct {
	planets []*Planet
}

func (executor *Executor) execMain(opts *Opts) {
	log.Debugf("Function: execMain")
	var wg sync.WaitGroup

	log.Debugln("pretty " + strconv.FormatBool(opts.Pretty))
	log.Debugln("script " + opts.ScriptName)
	log.Debugln("command " + opts.Command)
	for _, planet := range opts.Planets {
		log.Debugf("planet %s", planet)
	}

	wg.Add(len(executor.planets))

	for _, planet := range executor.planets {
		// to avoid closure over the value planet. seems odd but it is recommended
		planet := planet
		go func() {
			planet.execute(opts)
			wg.Done()
		}()
	}
	wg.Wait()
	formatAndPrint(executor.planets, opts)
}
