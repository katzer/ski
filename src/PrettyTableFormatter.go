package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
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
	prettyTableFormatter.addKey("Nr.")
	prettyTableFormatter.addKey("Planet-ID")
	prettyTableFormatter.addKey("Planet-Name")
	prettyTableFormatter.addKey("Planet-Address")
	prettyTableFormatter.addKey("Planet-Type")
}

func (prettyTableFormatter *PrettyTableFormatter) format(planet Planet, opts *Opts) {
	decodedJSON, err := decode(planet, planet.outputStruct.output)
	fullTable := prettyTableFormatter.addMetadata(planet)
	if err == nil {
		fullTable = prettyTableFormatter.normalizeTable(fullTable, decodedJSON)
	}
	set := Dataset{fullTable, nil}
	log.Debugf("prettyTableFormatter.format()")
	log.Debugf("prettyTableFormatter.keys %v", prettyTableFormatter.keys)
	log.Debugf("prettyTableFormatter.keys length %i", len(prettyTableFormatter.keys))
	log.Debugf("fullTable %v", fullTable)
	log.Debugf("fullTable length %i", len(fullTable))
	prettyTableFormatter.sets = append(prettyTableFormatter.sets, set)
}

func (prettyTableFormatter *PrettyTableFormatter) addMetadata(planet Planet) map[string]string {
	var toComplete = make(map[string]string)
	address := fmt.Sprintf("%s@%s", planet.user, planet.host)
	prettyTableFormatter.addEntry("Nr.", strconv.Itoa(planet.outputStruct.position), toComplete)
	prettyTableFormatter.addEntry("Planet-ID", planet.id, toComplete)
	prettyTableFormatter.addEntry("Planet-Name", planet.name, toComplete)
	prettyTableFormatter.addEntry("Planet-Address", address, toComplete)
	prettyTableFormatter.addEntry("Planet-Type", planet.planetType, toComplete)
	return toComplete
}

func (prettyTableFormatter *PrettyTableFormatter) normalizeTable(toReturn map[string]string, toNormalize [][]string) map[string]string {
	keys := toNormalize[0][:]
	log.Debugf("prettyTableFormatter.normalizeTable()")
	log.Debugf(".keys %v", keys)
	log.Debugf(".keys length %i", len(keys))
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
}

func (dataset *Dataset) makePrintView(keys map[int]string) {
	for i := 0; i <= len(keys)-1; i++ {
		if dataset.data[keys[i]] == "" {
			dataset.printView = append(dataset.printView, "-")
			continue
		}
		dataset.printView = append(dataset.printView, dataset.data[keys[i]])
	}
}

func (prettyTableFormatter *PrettyTableFormatter) cutMapToSlice(toCut map[string]bool) []string {
	toReturn := make([]string, 0)
	for i := 0; i < len(prettyTableFormatter.orderedKeys); i++ {
		toReturn = append(toReturn, prettyTableFormatter.orderedKeys[i])
	}
	return toReturn
}

func (prettyTableFormatter *PrettyTableFormatter) printTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetHeader(prettyTableFormatter.cutMapToSlice(prettyTableFormatter.keys))

	for _, set := range prettyTableFormatter.sets {
		table.Append(set.printView)
		log.Debugf("prettyTableFormatter.printTable()")
		log.Debugf("set.printView %v", set.printView)
		log.Debugf("set.printView length %i", len(set.printView))
	}

	table.Render() // Send output
}

func (prettyTableFormatter *PrettyTableFormatter) execute() {
	prettyTableFormatter.fillSets()
	prettyTableFormatter.printTable()
}
