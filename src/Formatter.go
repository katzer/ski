package main

import ()

//Formatter a struct remembering the different formatter
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
	if !opts.prettyFlag {
		return
	}
	if opts.template != "" {
		formatter.prettyTableFormatter.execute()
		return
	}
	formatter.prettyFormatter.execute()
}
