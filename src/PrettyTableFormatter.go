package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

const prettyPythonScriptName = "texttable.py"

// PTFAdapter ...
type PTFAdapter struct {
	real *PrettyTableFormatter
}

func (ptfAdapter PTFAdapter) init() {
	ptfAdapter.real.init()
}

func (ptfAdapter PTFAdapter) format(planets []Planet, opts *Opts, writer io.Writer) {
	ptfAdapter.real.format(planets, opts, writer)
}

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
	errored   bool
}

func (dataset *Dataset) makePrintView(keys map[int]string) {
	filler := "-"
	if dataset.errored {
		filler = makeRed("-")
	}
	for i := 0; i <= len(keys)-1; i++ {
		if dataset.data[keys[i]] == "" {
			dataset.printView = append(dataset.printView, filler)
			continue
		}
		dataset.printView = append(dataset.printView, dataset.data[keys[i]])
	}
}

func (prettyTableFormatter *PrettyTableFormatter) init() {
	prettyTableFormatter.keys = make(map[string]bool)
	prettyTableFormatter.orderedKeys = make(map[int]string)
	prettyTableFormatter.sets = make([]Dataset, 0)
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
	log.Debugf("len(prettyTableFormatter.orderedKeys) = %d\n", len(prettyTableFormatter.orderedKeys))
	prettyTableFormatter.addKey(key)
	if table[key] != "" {
		table[key] += ", " + value
		return
	}
	table[key] = value
}

func (prettyTableFormatter *PrettyTableFormatter) addKey(key string) {
	if !prettyTableFormatter.keys[key] {
		index := len(prettyTableFormatter.orderedKeys)
		log.Debugf("adding key : %s at [%d]\n", key, index)
		prettyTableFormatter.orderedKeys[index] = key
		prettyTableFormatter.keys[key] = true
	}
}

func (prettyTableFormatter *PrettyTableFormatter) fillSets() {
	for i, set := range prettyTableFormatter.sets {
		set.makePrintView(prettyTableFormatter.orderedKeys)
		prettyTableFormatter.sets[i] = set
	}
}

func (prettyTableFormatter *PrettyTableFormatter) cutMapToSlice(toCut map[string]bool) []string {
	toReturn := make([]string, 0)
	for i := 0; i < len(prettyTableFormatter.orderedKeys); i++ {
		toReturn = append(toReturn, prettyTableFormatter.orderedKeys[i])
	}
	return toReturn
}

func (prettyTableFormatter *PrettyTableFormatter) printTable(writer io.Writer) {

	table := tablewriter.NewWriter(writer)
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

func (prettyTableFormatter *PrettyTableFormatter) format(planets []Planet, opts *Opts, writer io.Writer) {
	tableFormatter := TableFormatter{} // TODO find out if it is necessary or get it from the factory
	for _, planet := range planets {
		jsonString := tableFormatter.formatPlanet(planet, opts)
		prettyTableFormatter.createSetForPlanet(jsonString, planet, opts)
	}
	prettyTableFormatter.fillSets()
	prettyTableFormatter.printTable(writer)
}

func (prettyTableFormatter *PrettyTableFormatter) createSetForPlanet(json string, planet Planet, opts *Opts) {
	table := make(map[string]string)
	number := strconv.Itoa(planet.outputStruct.position)
	address := fmt.Sprintf("%s@%s", planet.user, planet.host)
	id := planet.id
	name := planet.name
	planetType := planet.planetType
	if planet.outputStruct.errored || !planet.valid {
		number = makeRed(number)
		id = makeRed(id)
		name = makeRed(name)
		address = makeRed(address)
		planetType = makeRed(planetType)
	}
	prettyTableFormatter.addEntry("Nr.", number, table)
	prettyTableFormatter.addEntry("ID", id, table)
	prettyTableFormatter.addEntry("Name", name, table)
	prettyTableFormatter.addEntry("Address", address, table)
	prettyTableFormatter.addEntry("Type", planetType, table)
	var err error
	if planet.outputStruct.table, err = decode(json); err == nil {
		table = prettyTableFormatter.normalizeTable(table, planet.outputStruct.table)
	} else {
		log.Warnf("Error decoding json for planet %s", planet.id)
		planet.outputStruct.errored = true
		planet.outputStruct.output = json // TODO check if necessary
	}
	set := Dataset{table, nil, (!planet.valid || planet.outputStruct.errored)}
	prettyTableFormatter.sets = append(prettyTableFormatter.sets, set)
}
