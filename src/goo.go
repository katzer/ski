package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

/**
################################################################################
								Main-Section
################################################################################
*/

/**
*	StructuredOuput:
*	A thingy thing thing
 */
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

	prettyFlag, scriptFlag, scriptPath, command, planets, debugFlag, typeFlag, loadFlag := procArgs(args)

	outputList := make([]StructuredOuput, len(planets))

	_ = prettyFlag
	if debugFlag {
		fmt.Println(args)
		fmt.Println("prettyflag " + strconv.FormatBool(prettyFlag))
		fmt.Println("scriptflag " + strconv.FormatBool(scriptFlag))
		fmt.Println("scriptpath " + scriptPath)
		fmt.Println("command " + command)
		for _, planet := range planets {
			fmt.Println("planet " + planet)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(planets))
	for i, planet := range planets {
		if typeFlag {
			fmt.Println("The type of " + planet + " is " + getType(planet))
		}

		switch getType(planet) {
		case "server":
			connDet := getConnDet(planet)
			outputList[i].planet = planet
			if scriptFlag {
				go upAndExecSSHScript(connDet, scriptPath, &wg, &outputList[i], loadFlag)
			} else {
				go execSSHCommand(connDet, command, &wg, true, &outputList[i], loadFlag)
			}
		case "db":
			dbDet := "" //getDBDet(planet)
			outputList[i].planet = planet
			if scriptFlag {
				go upAndExecDBScript(dbDet, scriptPath, &wg, &outputList[i], loadFlag)
			} else {
				go execDBCommand(dbDet, command, &wg, true, &outputList[i], loadFlag)
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
	formatAndPrint(outputList, prettyFlag, scriptFlag, scriptPath, command)
}
