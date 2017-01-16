package main

import (
	"bytes"
	"fmt"
	"log"
)

var (
	buf    bytes.Buffer
	logger *log.Logger
)

func debugPrintOpts(opts *Opts) {
	debugString := ""
	debugString = fmt.Sprintf("%sopts:\n", debugString)
	debugString = fmt.Sprintf("%sprettyFlag: %t\n", debugString, opts.prettyFlag)
	debugString = fmt.Sprintf("%sscriptFlag: %t\n", debugString, opts.scriptFlag)
	debugString = fmt.Sprintf("%stypeFlag: %t\n", debugString, opts.typeFlag)
	debugString = fmt.Sprintf("%sdebugFlag: %t\n", debugString, opts.debugFlag)
	debugString = fmt.Sprintf("%sloadFlag: %t\n", debugString, opts.loadFlag)
	debugString = fmt.Sprintf("%shelpFlag: %t\n", debugString, opts.helpFlag)
	debugString = fmt.Sprintf("%sversionFlag: %t\n", debugString, opts.versionFlag)
	debugString = fmt.Sprintf("%stableFlag: %t\n", debugString, opts.tableFlag)
	debugString = fmt.Sprintf("%sscriptPath: %s\n", debugString, opts.scriptPath)
	debugString = fmt.Sprintf("%scommand: %s\n", debugString, opts.command)
	debugString = fmt.Sprintf("%splanets: %v\n", debugString, opts.planets)
	debugString = fmt.Sprintf("%splanetsCount: %d\n", debugString, opts.planetsCount)
	debugString = fmt.Sprintf("%scurrentDet: %s\n", debugString, opts.currentDet)
	debugString = fmt.Sprintf("%scurrentDBDet: %s\n", debugString, opts.currentDBDet)
	fmt.Print(debugString)
	log.Output(1, debugString)
}

func debugPrintStructuredOutput(strucOut *StructuredOuput) {
	debugString := ""
	debugString = fmt.Sprintf("%sstrucOut: %v\n", debugString, strucOut)
	debugString = fmt.Sprintf("%splanet: %s\n", debugString, strucOut.planet)
	debugString = fmt.Sprintf("%sout: %s\n", debugString, strucOut.output)
	debugString = fmt.Sprintf("%smaxLineLength: %d\n", debugString, strucOut.maxOutLength)
	fmt.Print(debugString)
	log.Output(1, debugString)

}

func debugPrintPlanets(planets []Planet) {
	debugString := ""
	for _, planet := range planets {
		debugString = fmt.Sprintf("%s%s\n", debugString, planet)
	}
	fmt.Print(debugString)
	log.Output(1, debugString)
}

func debugPrintString(message string) {
	debugString := ""
	debugString = fmt.Sprintf("%s", message)
	fmt.Print(debugString)
	log.Output(1, debugString)
}

func printDebugStart() {
	debugString := ""
	debugString = fmt.Sprintf("###################################### Program Start ######################################\n")
	fmt.Print(debugString)
	log.Output(1, debugString)
}

func printDebugEnd() {
	debugString := ""
	debugString = fmt.Sprintf("###################################### Program End ######################################\n")
	fmt.Print(debugString)
	log.Output(1, debugString)
}
