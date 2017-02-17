package main

import (
	"fmt"
)

// PrettyFormatter displays output from one or multiple planets in a neat, orderly fashion
type PrettyFormatter struct {
}

func (prettyFormatter *PrettyFormatter) format(toFormat StructuredOuput, counter int) string {
	var formatted string
	formatted = fmt.Sprintf("%d", counter)
	if counter/10 < 1 {
		formatted = fmt.Sprintf("%s ", formatted)
	}
	formatted = fmt.Sprintf("%s   ", formatted)
	formatted = fmt.Sprintf("%s%s", formatted, toFormat.planet)
	whitespace := printWhite(planetLength - len(toFormat.planet))
	formatted = fmt.Sprintf("%s%s", formatted, whitespace)
	indented := printIndented(toFormat.output, 26, true)
	formatted = fmt.Sprintf("%s%s", formatted, indented)
	return formatted
}
