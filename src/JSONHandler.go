package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// JSONReport - The true output of a deeply disturbed,
// incurably depressed and manically suicidal program named ski.
type JSONReport struct {
	Meta    Opts            `json:"meta,omitempty"`
	Planets []PlanetWrapper `json:"planets"`
}

// PlanetWrapper ...
type PlanetWrapper struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	User       string `json:"user"`
	Host       string `json:"host"`
	PlanetType string `json:"planet_type"`
	DbID       string `json:"db_id"`
	Valid      bool   `json:"valid"`
	Output     string `json:"output"`
	Index      int    `json:"index"`
	Errored    bool   `json:"errored"`
}

func decode(jsonObject string) ([][]string, error) {
	var jsonBlob = []byte(jsonObject)
	var toReturn = make([][]string, 0)
	err := json.Unmarshal(jsonBlob, &toReturn)
	if err != nil {
		log.Errorln(err)
		return make([][]string, 0), err
	}
	return toReturn, nil
}

func writeResultAsJSON(planets []Planet, opts *Opts, writer io.Writer) {
	allInOne := JSONReport{}
	allInOne.Planets = make([]PlanetWrapper, len(planets))
	allInOne.Meta = *opts
	for i, planet := range planets {
		wrapper := PlanetWrapper{
			ID:         planet.id,
			Name:       planet.name,
			User:       planet.user,
			Host:       planet.host,
			PlanetType: planet.planetType,
			DbID:       planet.dbID,
			Valid:      planet.valid,
			Output:     planet.outputStruct.output,
			Index:      i,
			Errored:    planet.outputStruct.errored,
		}
		allInOne.Planets[i] = wrapper
	}

	json, err := json.MarshalIndent(allInOne, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling the output as json %s\n", err.Error())
		fmt.Fprintf(os.Stderr, "Non json %v\n", allInOne)
		os.Exit(1)
	}

	fmt.Fprintf(writer, "%s\n", json)
}

// creates a task from a json file
func createATaskFromJobFile(jsonFile string) (opts Opts) {
	job := Opts{}
	wcopy := jsonFile // assumption abs path
	tokens := strings.Split(jsonFile, string(os.PathSeparator))

	if len(tokens) == 1 {
		// relative path given, read from jobs folder
		wcopy = path.Join(os.Getenv("ORBIT_HOME"), "jobs", jsonFile)
	}
	var err error
	var bytes []byte
	if bytes, err = ioutil.ReadFile(wcopy); err != nil {
		errorMessage := fmt.Sprintf("%s : %s", err.Error(), jsonFile)
		fmt.Fprint(os.Stderr, errorMessage)
		log.Fatal(errorMessage)
	}

	if json.Unmarshal(bytes, &job); err != nil {
		errorMessage := fmt.Sprintf("%s : %s", err.Error(), jsonFile)
		fmt.Fprint(os.Stderr, errorMessage)
		log.Fatal(errorMessage)
	}

	log.Debugf("Read a task from %s:", jsonFile)
	log.Debugf("Unmarshalled %v", job)
	return job
}
