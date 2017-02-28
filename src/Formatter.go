package main

import (
	"strings"
)

//Formatter a struct remebering the different formatter
type Formatter struct {
	prettyFormatter      PrettyFormatter
	tableFormatter       TableFormatter
	prettyTableFormatter PrettyTableFormatter
}

func (formatter *Formatter) init() {
	formatter.prettyFormatter = PrettyFormatter{}
	formatter.tableFormatter = TableFormatter{}
	formatter.prettyTableFormatter = PrettyTableFormatter{}

	formatter.prettyTableFormatter.init()
}

func (formatter *Formatter) format(toPrint StructuredOuput, counter int, opts *Opts) string {
	var formatted string
	if !opts.pretty {
		if opts.template != "" {
			formatted = formatter.tableFormatter.format(toPrint.output, opts)
		} else {
			formatted = toPrint.output
		}
	} else {
		if opts.template != "" {
			preFormatted := formatter.tableFormatter.format(toPrint.output, opts)
			formatter.prettyTableFormatter.format(preFormatted, opts)
		} else {
			formatted = formatter.prettyFormatter.format(toPrint, counter)
		}
	}
	return formatted
}

func parseFSMOutput(toParse string) map[string]string {
	var parsed map[string]string
	parsed = make(map[string]string)
	split := strings.Split(toParse, "\n")
	split = split[0 : len(split)-2]
	for _, entry := range split {
		row := strings.Split(entry, " ")
		parsed[strings.TrimSuffix(row[0], ",")] = strings.Join(row[1:], "")
	}
	return parsed
}
