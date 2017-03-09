package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

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
	prettyFormatter.addKey("Nr.")
	prettyFormatter.addKey("Planet-ID")
	prettyFormatter.addKey("Planet-Name")
	prettyFormatter.addKey("Planet-Address")
	prettyFormatter.addKey("Planet-Type")
}

func (prettyFormatter *PrettyFormatter) addMetadata(toComplete map[string]string, planet *Planet) map[string]string {
	address := fmt.Sprintf("%s@%s", planet.user, planet.host)
	prettyFormatter.addEntry("Nr.", strconv.Itoa(planet.outputStruct.position), toComplete)
	prettyFormatter.addEntry("Planet-ID", planet.id, toComplete)
	prettyFormatter.addEntry("Planet-Name", planet.name, toComplete)
	prettyFormatter.addEntry("Planet-Address", address, toComplete)
	prettyFormatter.addEntry("Planet-Type", planet.planetType, toComplete)
	return toComplete
}

func (prettyFormatter *PrettyFormatter) addEntry(key string, value string, table map[string]string) {
	prettyFormatter.addKey(key)
	if table[key] != "" {
		table[key] += ", " + value
		return
	}
	table[key] = value
}

func (prettyFormatter *PrettyFormatter) addKey(key string) {
	if !prettyFormatter.keys[key] {
		prettyFormatter.orderedKeys[len(prettyFormatter.orderedKeys)] = key
		prettyFormatter.keys[key] = true
	}
}

func (prettyFormatter *PrettyFormatter) format(planet *Planet) {
	var completeTable = make(map[string]string)
	prettyFormatter.addMetadata(completeTable, planet)
	prettyFormatter.addEntry("output", planet.outputStruct.output, completeTable)
	set := Dataset{completeTable, nil}
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

func (prettyFormatter *PrettyFormatter) printTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetHeader(prettyFormatter.cutMapToSlice(prettyFormatter.keys))

	for _, set := range prettyFormatter.sets {
		log.Debugf("prettyTable.printTable settoappend: %v\n", set.printView)
		table.Append(set.printView)
	}

	table.Render() // Send output
}

func (prettyFormatter *PrettyFormatter) execute() {
	prettyFormatter.fillSets()
	prettyFormatter.printTable()
}
