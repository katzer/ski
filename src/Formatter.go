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
	if !opts.Pretty {
		if opts.Template != "" {
			formatted = formatter.tableFormatter.format(planet, opts)
		} else {
			formatted = planet.outputStruct.output
		}
	} else {
		if opts.Template != "" {
			planet.outputStruct.output = formatter.tableFormatter.format(planet, opts)
			formatter.prettyTableFormatter.format(planet, opts)
		} else {
			formatter.prettyFormatter.format(planet)
		}
	}
	return formatted
}

func (formatter *Formatter) execute(opts *Opts) {
	if !opts.Pretty {
		return
	}
	if opts.Template != "" {
		formatter.prettyTableFormatter.execute()
		return
	}
	formatter.prettyFormatter.execute()
}
