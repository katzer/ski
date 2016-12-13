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
*	Main function
 */
func main() {

	var args []string = os.Args

	prettyFlag, scriptFlag, scriptPath, command, planets, debugFlag, typeFlag := procArgs(args)

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
	for _, planet := range planets {
		if typeFlag {
			fmt.Println("The type of " + planet + " is " + getType(planet))
		}
		switch getType(planet) {
		case "server":
			var connDet string = getConnDet(planet)
			if scriptFlag {
				go upAndExecScript(connDet, scriptPath, &wg, prettyFlag)
			} else {
				go execCommand(connDet, command, &wg, true, prettyFlag)
			}
		case "db":
			fmt.Fprintf(os.Stderr, "This Type of Connection is not yet supported.")
			os.Exit(1)
		case "web":
			fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
			os.Exit(1)
		default:
			wg.Done()
		}
	}
	wg.Wait()
}
