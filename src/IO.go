package main

import (
	"bytes"
	"fmt"
	"strings"
)

const planetLength int = 21

/**
*	Prints the given String indented by the given spaces.
 */
func printIndented(msg string, indent int, exceptFirst bool) {
	charString := []rune(msg)
	var buffer bytes.Buffer
	_ = exceptFirst
	toAppend := ""
	if !exceptFirst {
		for i := 0; i < indent; i++ {
			buffer.WriteByte(32)
		}
	}
	for _, char := range charString {
		if char == 10 {
			fmt.Println(buffer.String())
			buffer.Reset()
			for i := 0; i < indent; i++ {
				buffer.WriteByte(32)
			}
		}
		if char != 10 {
			toAppend = fmt.Sprintf("%c", char)
			buffer.WriteString(toAppend)
		}
	}
	fmt.Println(buffer.String())
}

func printHeadline(scriptFlag bool, scriptName string, command string, indent int) {
	fmt.Print("NR   PLANET               ")
	if scriptFlag {
		printIndented(scriptName, indent, true)
	} else {
		printIndented(command, indent, true)
	}
	fmt.Println("================================================================================")
}

func printWhite(length int) {
	for i := 0; i < length; i++ {
		fmt.Print(" ")
	}
}

func formatAndPrint(toPrint []StructuredOuput, opts *Opts) {
	if opts.prettyFlag && !opts.tableFlag {
		printHeadline(opts.scriptFlag, opts.scriptName, opts.command, 80)
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
