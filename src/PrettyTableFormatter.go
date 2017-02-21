package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

const prettyPythonScriptName = "texttable.py"

// PrettyTableFormatter prints input in tabular format
type PrettyTableFormatter struct {
	keys        map[string]bool
	orderedKeys map[int]string
	sets        []Dataset
}

// Dataset : One set of data
type Dataset struct {
	data      map[string]string
	printView []string
}

func (prettyTableFormatter *PrettyTableFormatter) init() {
	prettyTableFormatter.keys = make(map[string]bool)
	prettyTableFormatter.orderedKeys = make(map[int]string)
	prettyTableFormatter.sets = make([]Dataset, 0)
}

func (prettyTableFormatter *PrettyTableFormatter) format(toFormat string, opts *Opts) {
	var decodedJSON = make([][]string, 0)
	decodedJSON = decode(toFormat)
	normalizedTable := prettyTableFormatter.normalizeTable(decodedJSON)
	var set Dataset
	set.data = normalizedTable
	prettyTableFormatter.sets = append(prettyTableFormatter.sets, set)

}

func (prettyTableFormatter *PrettyTableFormatter) normalizeTable(toNormalize [][]string) map[string]string {
	var toReturn = make(map[string]string)
	keys := toNormalize[0][:]
	values := toNormalize[1:][:]
	skip := false
	for _, entry := range values {
		for i, value := range entry {
			if skip {
				skip = false
				continue
			}
			if value == "" {
				continue
			}
			if strings.Contains(keys[i], "Key") {
				prettyTableFormatter.addEntry(value, entry[i+1], toReturn)
				skip = true
				continue
			}
			prettyTableFormatter.addEntry(keys[i], value, toReturn)
		}
	}
	return toReturn
}

func (prettyTableFormatter *PrettyTableFormatter) addEntry(key string, value string, table map[string]string) {
	prettyTableFormatter.addKey(key)
	if table[key] != "" {
		table[key] += ", " + value
		return
	}
	table[key] = value
}

func (prettyTableFormatter *PrettyTableFormatter) addKey(key string) {
	if !prettyTableFormatter.keys[key] {
		prettyTableFormatter.orderedKeys[len(prettyTableFormatter.orderedKeys)] = key
		prettyTableFormatter.keys[key] = true
	}
}

func (prettyTableFormatter *PrettyTableFormatter) fillSets() {
	for i, set := range prettyTableFormatter.sets {
		set.makePrintView(prettyTableFormatter.orderedKeys)
		prettyTableFormatter.sets[i] = set
	}
	log.Debugf("prettyTableFormatter.fillSets dataset.printview: %v\n", prettyTableFormatter.sets[0].printView)
}

func (dataset *Dataset) makePrintView(keys map[int]string) {
	for i := 0; i < len(keys)-1; i++ {
		if dataset.data[keys[i]] == "" {
			dataset.printView = append(dataset.printView, "-")
			continue
		}
		dataset.printView = append(dataset.printView, dataset.data[keys[i]])
	}
	log.Debugf("prettyTableFormatter.makePrintView keys: %v\n", keys)
	log.Debugf("prettyTableFormatter.makePrintView dataset.data: %v\n", dataset.data)
	log.Debugf("prettyTableFormatter.makePrintView dataset.data len: %d\n", len(dataset.data))
	log.Debugf("prettyTableFormatter.makePrintView dataset.printview: %v\n", dataset.printView)
}

func (prettyTableFormatter *PrettyTableFormatter) cutMapToSlice(toCut map[string]bool) []string {
	toReturn := make([]string, 0)
	for i := 0; i < len(prettyTableFormatter.orderedKeys)-1; i++ {
		toReturn = append(toReturn, prettyTableFormatter.orderedKeys[i])
	}
	log.Debugf("prettyTableFormatter.cutMapToSlice: %v\n", toReturn)
	return toReturn
}

func (prettyTableFormatter *PrettyTableFormatter) printTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(prettyTableFormatter.cutMapToSlice(prettyTableFormatter.keys))

	for _, set := range prettyTableFormatter.sets {
		log.Debugf("prettyTable.printTable settoappend: %v\n", set.printView)
		table.Append(set.printView)
	}

	table.Render() // Send output
}

func (prettyTableFormatter *PrettyTableFormatter) execute() {
	prettyTableFormatter.fillSets()
	prettyTableFormatter.printTable()
}
