package main

import (
	"fmt"
	"io"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

// PFAdapter ...
type PFAdapter struct {
	real *PrettyFormatter
}

func (pfAdapter PFAdapter) init() {
	pfAdapter.real.init()
}

func (pfAdapter PFAdapter) format(planets []Planet, opts *Opts, writer io.Writer) {
	pfAdapter.real.format(planets, opts, writer)
}

// PrettyFormatter displays output from one or multiple planets in a neat, orderly fashion
type PrettyFormatter struct {
	keys        map[string]bool
	orderedKeys map[int]string
	sets        []Dataset
}

func (prettyFormatter *PrettyFormatter) init() {
	prettyFormatter.keys = make(map[string]bool)
	prettyFormatter.orderedKeys = make(map[int]string)
	prettyFormatter.sets = make([]Dataset, 0)
}

func (prettyFormatter *PrettyFormatter) addEntry(key string, value string, table map[string]string) {
	log.Debugf("len(prettyFormatter.orderedKeys) = %d\n", len(prettyFormatter.orderedKeys))
	prettyFormatter.addKey(key)
	if table[key] != "" {
		table[key] += ", " + value
		return
	}
	table[key] = value
}

func (prettyFormatter *PrettyFormatter) addKey(key string) {
	if !prettyFormatter.keys[key] {
		index := len(prettyFormatter.orderedKeys)
		log.Debugf("adding key : %s at [%d]\n", key, index)
		prettyFormatter.orderedKeys[index] = key
		prettyFormatter.keys[key] = true
	}
}

func (prettyFormatter *PrettyFormatter) createSetForPlanet(planet Planet) {
	var completeTable = make(map[string]string)
	id := planet.id
	name := planet.name
	planetType := planet.planetType
	address := fmt.Sprintf("%s@%s", planet.user, planet.host)
	number := strconv.Itoa(planet.outputStruct.position)
	output := planet.outputStruct.output
	if planet.outputStruct.errored || !planet.valid {
		number = makeRed(number)
		id = makeRed(id)
		name = makeRed(name)
		address = makeRed(address)
		planetType = makeRed(planetType)
		output = makeRed(planet.outputStruct.output)
	}
	prettyFormatter.addEntry("Nr.", number, completeTable)
	prettyFormatter.addEntry("ID", id, completeTable)
	prettyFormatter.addEntry("Name", name, completeTable)
	prettyFormatter.addEntry("Address", address, completeTable)
	prettyFormatter.addEntry("Type", planetType, completeTable)
	prettyFormatter.addEntry("output", output, completeTable)

	set := Dataset{completeTable, nil, (planet.valid && !planet.outputStruct.errored)}
	prettyFormatter.sets = append(prettyFormatter.sets, set)
}

func (prettyFormatter *PrettyFormatter) fillSets() {
	for i, set := range prettyFormatter.sets {
		set.makePrintView(prettyFormatter.orderedKeys)
		prettyFormatter.sets[i] = set
	}
	log.Debugf("prettyFormatter.fillSets dataset.printview: %v\n", prettyFormatter.sets[0].printView)
}

func (prettyFormatter *PrettyFormatter) cutMapToSlice(toCut map[string]bool) []string {
	toReturn := make([]string, 0)
	for i := 0; i < len(prettyFormatter.orderedKeys); i++ {
		toReturn = append(toReturn, prettyFormatter.orderedKeys[i])
	}
	log.Debugf("prettyFormatter.cutMapToSlice: %v\n", toReturn)
	return toReturn
}

func (prettyFormatter *PrettyFormatter) printTable(writer io.Writer) {

	table := tablewriter.NewWriter(writer)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetHeader(prettyFormatter.cutMapToSlice(prettyFormatter.keys))
	table.SetAutoWrapText(false)

	for _, set := range prettyFormatter.sets {
		log.Debugf("prettyTable.printTable settoappend: %v\n", set.printView)
		table.Append(set.printView)
	}

	table.Render() // Send output
}

func (prettyFormatter *PrettyFormatter) format(planets []Planet, opts *Opts, writer io.Writer) {
	log.Debugf("planets: %v \n", planets)
	log.Debugf("opts : %s", opts.String())
	for _, planet := range planets {
		prettyFormatter.createSetForPlanet(planet)
	}
	prettyFormatter.fillSets()
	prettyFormatter.printTable(writer)
}
