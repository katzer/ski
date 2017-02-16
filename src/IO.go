package main

import (
	"fmt"
	"strings"
)

const planetLength int = 21

/**
*	Prints the given String indented by the given spaces.
 */
func printIndented(msg string, indent int, exceptFirst bool) string {
	whitespace := printWhite(indent)
	rows := strings.SplitAfter(msg, "\n")
	indented := ""
	for count, row := range rows {
		if count == 0 && exceptFirst {
			indented = row
			continue
		}
		if count == len(rows)-1 {
			continue
		}
		indented = fmt.Sprintf("%s%s%s", indented, whitespace, row)
	}
	return indented
}

func printHeadline(indent int, opts *Opts) {
	headline := "NR   PLANET               "

	if opts.scriptFlag {
		headline = fmt.Sprintf("%s%s\n", headline, printIndented(opts.scriptName, indent, true))
	} else {
		headline = fmt.Sprintf("%s%s\n", headline, printIndented(opts.command, indent, true))
	}
	separator := "================================================================================"
	headline = fmt.Sprintf("%s%s\n", headline, separator)
	fmt.Print(headline)
}

func printWhite(length int) string {
	whitespace := ""
	for i := 0; i < length; i++ {
		whitespace = fmt.Sprintf("%s ", whitespace)
	}
	return whitespace
}

func formatAndPrint(toPrint []StructuredOuput, opts *Opts) {
	if opts.prettyFlag && !opts.tableFlag {
		printHeadline(80, opts)
	}
	for i, entry := range toPrint {
		formatted := format(entry, i, opts)

		fmt.Print(formatted)
	}
}

func trimDBMetaInformations(strucOut *StructuredOuput) {
	cleaned := strings.Split(strucOut.output, "\n")
	strucOut.output = strings.Join(cleaned[:len(cleaned)-3], "")
}
