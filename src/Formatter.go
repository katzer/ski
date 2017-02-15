package main

import (
	"fmt"
	"strconv"
)

func format(toPrint StructuredOuput, counter int, opts *Opts) string {
	var formatted string
	if !opts.prettyFlag {
		if opts.tableFlag {
			var formatter TableFormatter
			formatted = formatter.format(toPrint.output, opts)
		} else {
			formatted = toPrint.output
		}
	} else {
		if opts.tableFlag {
			var preFormatter TableFormatter
			preFormatted := preFormatter.format(toPrint.output, opts)
			var formatter PrettyTableFormatter
			formatted = formatter.format(preFormatted, opts)
		} else {
			fmt.Print(strconv.Itoa(counter) + "")
			if counter/10 < 1 {
				fmt.Print(" ")
			}
			if counter/100 < 1 {
				fmt.Print(" ")
			}
			if counter/1000 < 1 {
				fmt.Print(" ")
			}
			fmt.Print(" ")
			fmt.Print(toPrint.planet)
			printWhite(planetLength - len(toPrint.planet))
			printIndented(toPrint.output, 26, true)
		}
	}
	return formatted
}
