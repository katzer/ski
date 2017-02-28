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
	formatter.prettyFormatter.init()
}

func (formatter *Formatter) format(planet Planet, opts *Opts) string {
	var formatted string
	if !opts.prettyFlag {
		if opts.template != "" {
			formatted = formatter.tableFormatter.format(planet, opts)
		} else {
			formatted = planet.outputStruct.output
		}
	} else {
		if opts.template != "" {
			planet.outputStruct.output = formatter.tableFormatter.format(planet, opts)
			formatter.prettyTableFormatter.format(planet, opts)
		} else {
			formatter.prettyFormatter.format(planet)
		}
	}
	return formatted
}

func (formatter *Formatter) execute(opts *Opts) {
	if opts.prettyFlag {
		if opts.template != "" {
			formatter.prettyTableFormatter.execute()
			return
		}
		formatter.prettyFormatter.execute()
	}
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
