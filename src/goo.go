package main

import (
	"fmt"
	"log"
	"os"
)

// StructuredOuput ...
type StructuredOuput struct {
	planet       string
	output       string
	maxOutLength int
}

const server = "server"
const webServer = "web"
const database = "db"

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
	// if this was moved to any function other than main, the logger wouldn't function as the file it writes to would be closed after the init function was finished
	if opts.debugFlag {
		// open a file
		f, err := os.OpenFile("testlog.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}

		// don't forget to close it
		defer f.Close()

		// Output to stderr instead of stdout, could also be a file.
		log.SetOutput(f)

		printDebugStart()
	}

	exec := makeExecutor(&opts)

	exec.execMain(&opts)

	if opts.debugFlag {
		printDebugEnd()
	}

}

func makeExecutor(opts *Opts) Executor {
	executor := Executor{}
	var planet Planet
	for _, entry := range opts.planets {
		var user string
		var host string
		var dbID string
		connDet := getPlanetDetails(entry)
		planet.outputStruct.planet = entry
		id := entry
		planetType := getType(entry)
		switch planetType {
		case server:
			user = getUser(connDet)
			host = getHost(connDet)
			dbID = ""
		case database:
			dbID, connDet = procDBDets(connDet)
			user = getUser(connDet)
			host = getHost(connDet)
		case webServer:
			fmt.Println("Usage of goo with web servers is not implemented")
			continue
		default:
			fmt.Printf("Unkown Type of target %s: %s\n", entry, planet.planetType)
			continue
		}
		executor.planets = append(executor.planets, Planet{id, user, host, planetType, dbID, StructuredOuput{id, "", 0}})

	}
	return executor
}

/**
*	Prints the current Version of the goo application
 */
func printVersion() {
	os.Stdout.WriteString("0.9")
}

/**
*	Prints the help dialog
 */
func printHelp() {
	fmt.Println(`usage: goo [options...] <planet>... -c="<command>"`)
	fmt.Println(`Options:`)
	fmt.Println(`-s="<path/to/script>", --script="<path/to/script>"  Execute script and return result`)
	fmt.Println(`-p, --pretty     Pretty print output as a table`)
	fmt.Println(`-t, --type       Show type of planet`)
	fmt.Println(`-h, --help       This help text`)
	fmt.Println(`-v, --version    Show version number`)
	fmt.Println(`-d, --debug	  Show extended debug informations`)
}
