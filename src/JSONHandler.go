package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

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

func createJSONReport(options map[string]string, planets []Planet, opts *Opts) {
	basename := strings.Split(options["job_name"], ".")[0]
	home := options["orbit_home"]
	reports := options["output"]
	if len(basename) == 0 || len(home) == 0 || len(reports) == 0 {
		log.Fatalf("Could not create json output for the job %s", opts.String())
	}

	folders := path.Join(home, reports, basename)
	err := os.MkdirAll(folders, 0744)

	if err != nil {
		log.Fatalf("Could not create json output for the job %s", opts.String())
	}

	now := time.Now()
	format := "%d-%02d-%02dT%02d_%02d_%02d"
	stamp := fmt.Sprintf(format, now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fileToWrite := strings.Join([]string{stamp, "json"}, ".")
	toCreate := path.Join(folders, fileToWrite)

	if writer, err := os.Create(toCreate); err == nil {
		defer writer.Close()
		writeResultAsJSON(planets, opts, writer)
		removeOldOutput(folders, opts.MaxToKeep)
		return
	}
	log.Fatalf("Could not create json output for the job %s", basename)
}

func removeOldOutput(dir string, maxToKeep int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Could not read files in folder %s", dir)
	}

	total := len(files)
	log.Infof("removeOldOutput:: Total logs: %d, max # to keep :%d", total, maxToKeep)
	if total > maxToKeep {
		diff := total - maxToKeep
		toDelete := files[:diff]

		for _, file := range toDelete {
			name := file.Name()
			abs := path.Join(dir, name)
			log.Infoln("removing old output " + abs)
			if err := os.Remove(abs); err != nil {
				log.Errorln("removing old output " + abs + " failed")
			}
		}
	}
}

// creates a task from a json file
func createATaskFromJobFile(jsonFile string) (opts Opts) {
	job := Opts{
		MaxToKeep: 10,
	}
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
