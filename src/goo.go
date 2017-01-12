package main

import (
	"fmt"
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

	opts.procArgs(args)
	if opts.helpFlag {
		printHelp()
	}
	if opts.versionFlag {
		printVersion()
	}

	exec := makeExecutor(&opts)

	exec.execMain(&opts)

}

func makeExecutor(opts *Opts) Executor {
	executor := Executor{}
	var planet Planet
	for _, entry := range opts.planets {
		planet = Planet{}
		connDet := getConnDet(entry)
		planet.outputStruct.planet = entry
		planet.id = entry
		planet.planetType = getType(entry)
		switch planet.planetType {
		case "server":
			planet.user = getUser(connDet)
			planet.host = getHost(connDet)
		case "db":
			var dbID string
			dbID, connDet = procDBDets(connDet)
			planet.user = getUser(connDet)
			planet.host = getHost(connDet)
			planet.dbID = dbID
		case "web":
			fmt.Println("Usage of goo with web servers is not implemented")
			continue
		default:
			fmt.Printf("Unkown Type of target %s: %s\n", entry, planet.planetType)
			continue
		}
		executor.planets = append(executor.planets, planet)

	}
	return executor
}
