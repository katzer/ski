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

	log "github.com/sirupsen/logrus"
)

// JSONReport - The true output of a deeply disturbed,
// incurably depressed and manically suicidal program named ski.
type JSONReport struct {
	Meta    Opts            `json:"meta,omitempty"`
	Planets []PlanetWrapper `json:"planets"`
}

// PlanetWrapper ...
type PlanetWrapper struct {
	ID        string `json:"id"`
	Keys      string `json:"keys"`
	Output    string `json:"output"`
	Errored   bool   `json:"errored"`
	CreatedAt string `json:"created_at"`
}

func decode(jsonObject string) ([][]string, error) {
	var jsonBlob = []byte(jsonObject)
	var toReturn = make([][]string, 0)
	err := json.Unmarshal(jsonBlob, &toReturn)
	if err != nil {
		log.Errorln("In JSONHandler:decode : ")
		log.Errorln("JSON String is " + jsonObject)
		log.Errorln(err)
		return make([][]string, 0), err
	}
	log.Debugf("Return is  %o:", toReturn)
	return toReturn, nil
}

func writeResultAsJSON(planets []Planet, opts *Opts, writer io.Writer) {
	allInOne := JSONReport{}
	allInOne.Planets = make([]PlanetWrapper, len(planets))
	allInOne.Meta = *opts
	for i, planet := range planets {
		wrapper := PlanetWrapper{
			ID:      planet.id,
			Keys:    strings.Join(planet.outputStruct.keys, ", "),
			Output:  planet.outputStruct.output,
			Errored: planet.outputStruct.errored,
			// RFC3339 is a subset of ISO 8601
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
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

	// If no pretty printing is invovled but a template is given
	// set the output of the tableformatter as the output of the job.
	if opts.Pretty == false && len(strings.TrimSpace(opts.Template)) > 0 {
		tableFormatter := TableFormatter{}
		for _, planet := range planets {
			jsonString := tableFormatter.formatPlanet(planet, opts)
			planet.outputStruct.output = jsonString
		}
	}

	now := time.Now()
	format := "%d-%02d-%02dT%02d_%02d_%02d"
	stamp := fmt.Sprintf(format, now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	toCreate := path.Join(folders, fmt.Sprintf(`%s%s`, stamp, jobExt))

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
	parentDir, basename := path.Split(jsonFile)

	if len(parentDir) == 0 {
		// relative path given, read from jobs folder
		wcopy = path.Join(os.Getenv("ORBIT_HOME"), "jobs", jsonFile)
	}

	tokens := strings.Split(basename, ".")
	if len(tokens) == 1 {
		wcopy = fmt.Sprintf(`%s%s`, wcopy, jobExt)
	}

	var err error
	var bytes []byte
	if bytes, err = ioutil.ReadFile(wcopy); err != nil {
		msg := fmt.Sprintf("%s : %s", err.Error(), jsonFile)
		fmt.Fprint(os.Stderr, msg)
		log.Fatal(msg)
	}

	if json.Unmarshal(bytes, &job); err != nil {
		msg := fmt.Sprintf("%s : %s", err.Error(), jsonFile)
		fmt.Fprint(os.Stderr, msg)
		log.Fatal(msg)
	}

	log.Debugf("Read a task from %s:", jsonFile)
	log.Debugf("Unmarshalled %v", job)
	return job
}
