package main

import (
	"fmt"
	"math"
	"strings"
)

const prettyPythonScriptName = "texttable.py"

type PrettyTableFormatter struct {
}

type Entry struct {
	key   string
	value string
}

func (prettyTableFormatter *PrettyTableFormatter) format(toFormat string, opts *Opts) string {
	var decodedJSON = make([][]string, 0)
	decodedJSON = decode(toFormat)
	normalizedTable := prettyTableFormatter.normalizeTable(decodedJSON)
	keys := ""
	seperator := ""
	values := ""
	for _, entry := range normalizedTable {
		maxLength := int(math.Max(float64(len(entry.key)), float64(len(entry.value))))
		keys = fmt.Sprintf("%s| %s%s ", keys, entry.key, strings.Repeat(" ", maxLength-len(entry.key)))
		values = fmt.Sprintf("%s| %s%s ", values, entry.value, strings.Repeat(" ", maxLength-len(entry.value)))
		seperator = fmt.Sprintf("%s--%s-", seperator, strings.Repeat("-", maxLength))
	}
	keys = fmt.Sprintf("%s|", keys)
	values = fmt.Sprintf("%s|", values)
	seperator = fmt.Sprintf("%s-", seperator)
	return fmt.Sprintf("%s\n%s\n%s\n", keys, seperator, values)

}

func (prettyTableFormatter *PrettyTableFormatter) parseFSMOutput(toParse string) map[string]string {
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

func (prettyTableFormatter *PrettyTableFormatter) normalizeTable(toNormalize [][]string) []Entry {
	var toReturn = make([]Entry, 0)
	keys := toNormalize[0][:]
	values := toNormalize[1:][:]
	skip := false
	for _, entry := range values {
		for i, value := range entry {
			if skip {
				skip = false
				continue
			} else if value != "" {
				if strings.Contains(keys[i], "Key") {
					toReturn = append(toReturn, Entry{value, entry[i+1]})
					skip = true
				} else if !strings.Contains(keys[i], "Value") {
					toReturn = append(toReturn, Entry{keys[i], value})
				}
			}
		}
	}
	return toReturn
}
