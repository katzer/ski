package main

import (
	"fmt"
)

func printOptsDebug(opts *Opts) {
	fmt.Printf("opts:\n")
	fmt.Printf("prettyFlag: %t\n", opts.prettyFlag)
	fmt.Printf("scriptFlag: %t\n", opts.scriptFlag)
	fmt.Printf("typeFlag: %t\n", opts.typeFlag)
	fmt.Printf("debugFlag: %t\n", opts.debugFlag)
	fmt.Printf("loadFlag: %t\n", opts.loadFlag)
	fmt.Printf("helpFlag: %t\n", opts.helpFlag)
	fmt.Printf("versionFlag: %t\n", opts.versionFlag)
	fmt.Printf("tableFlag: %t\n", opts.tableFlag)
	fmt.Printf("scriptPath: %s\n", opts.scriptPath)
	fmt.Printf("command: %s\n", opts.command)
	fmt.Printf("planets: %v\n", opts.planets)
	fmt.Printf("planetsCount: %d\n", opts.planetsCount)
	fmt.Printf("currentDet: %s\n", opts.currentDet)
	fmt.Printf("currentDBDet: %s\n", opts.currentDBDet)
}

func printStructuredOutputDebug(strucOut *StructuredOuput) {

	fmt.Printf("strucOut: %v\n", strucOut)
	fmt.Printf("planet: %s\n", strucOut.planet)
	fmt.Printf("out: %s\n", strucOut.output)
	fmt.Printf("maxLineLength", strucOut.maxOutLength)

}

func printDebugString(message string) {
	fmt.Printf(message)
}
