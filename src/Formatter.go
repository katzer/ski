package main

import (
	"strings"
)

func format(toPrint StructuredOuput, counter int, opts *Opts) string {
	var formatted string
	if !opts.prettyFlag {
		if opts.template != "" {
			var formatter TableFormatter
			formatted = formatter.format(toPrint.output, opts)
		} else {
			formatted = toPrint.output
		}
	} else {
		if opts.template != "" {
			var preFormatter TableFormatter
			preFormatted := preFormatter.format(toPrint.output, opts)
			var formatter PrettyTableFormatter
			formatted = formatter.format(preFormatted, opts)
		} else {
			var formatter PrettyFormatter
			formatted = formatter.format(toPrint, counter)
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
