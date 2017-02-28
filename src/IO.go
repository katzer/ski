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

	if !(opts.scriptName == "") {
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
	formatter := Formatter{}
	formatter.init()
	var formatted string
	if opts.pretty && opts.template == "" {
		printHeadline(80, opts)
	}
	for i, entry := range toPrint {
		formatted = formatter.format(entry, i, opts)
		fmt.Print(formatted)

	}
	if opts.pretty && opts.template != "" {
		formatter.prettyTableFormatter.execute()
	}
}

func trimDBMetaInformations(strucOut *StructuredOuput) {
	cleaned := strings.Split(strucOut.output, "\n")
	strucOut.output = strings.Join(cleaned[:len(cleaned)-3], "")
}

/**
*	Prepends the profile loading command and seperator to a command
 */
func makeLoadCommand(command string, opts *Opts) string {
	if opts.load {
		return fmt.Sprintf(`sh -lc "echo -----APPPLANT-ORBIT----- && %s "`, command)
	}
	return command
}

/**
*	Removes the output provided by the profile loading
 */
func cleanProfileLoadedOutput(output string, opts *Opts) string {
	if opts.load {
		splitOut := strings.Split(output, "-----APPPLANT-ORBIT-----\n")
		return splitOut[len(splitOut)-1]
	}
	return output
}
