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

type StructuredOuput struct{
	planet	string
	output	string
}

/**
*	Main function
 */
func main() {

	var args []string = os.Args

	prettyFlag, scriptFlag, scriptPath, command, planets, debugFlag, typeFlag := procArgs(args)


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
		if (typeFlag) {
			fmt.Println("The type of " + planet + " is " + getType(planet))
		}

		switch getType(planet) {
			case "server":
				var connDet string = getConnDet(planet)
				if(prettyFlag){
					fmt.Print("     " + planet)
				}
				outputList[i].planet = planet
				if scriptFlag {
					go upAndExecScript(connDet, scriptPath, &wg, &outputList[i])
				} else {
					go execCommand(connDet, command, &wg, true, &outputList[i])
				}
			case "db":
				fmt.Fprintf(os.Stderr, "This Type of Connection is not yet supported.")
				os.Exit(1)
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
