package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	victim := "[[\"1\"], [\"2\"]]"
	decoded, err := decode(victim)
	if err != nil {
		t.Fail()
	}
	for i, str := range decoded {
		fmt.Printf("decoded[%d] %v\n", i, str)
	}
	if decoded[0][0] != "1" {
		t.Fail()
	}
	if decoded[1][0] != "2" {
		t.Fail()
	}
}

func TestWriteResultAsJSON(t *testing.T) {
	var outDir string
	var err error
	jobFile := "job.js"
	// Setup
	name := "reports"
	backup := os.Getenv("ORBIT_HOME")
	os.Setenv("ORBIT_HOME", os.TempDir())

	if outDir, err = ioutil.TempDir(os.TempDir(), name); err != nil {
		t.Fail()
	}

	planet := Planet{
		id:         "app",
		name:       "App-Package 1",
		user:       "none",
		host:       "localhost",
		planetType: "server",
		dbID:       "",
		valid:      true,
		outputStruct: &StructuredOuput{
			planet:   "app",
			output:   "exit status 2",
			table:    nil,
			position: 0,
			errored:  true,
		},
	}
	planets := []Planet{planet}

	// "-j=/tmp/job.js", "-t=\"perlver_template\"", "-p", "-d=true", "app"
	opts := Opts{
		Template: "perlver_template",
		Pretty:   true,
		Debug:    true,
		Planets:  []string{"app"},
	}

	basename := path.Base(jobFile)
	fileToWrite := path.Join(outDir, basename)
	if writer, err := os.Create(fileToWrite); err != nil {
		t.Fail()
	} else {
		defer writer.Close()
		writeResultAsJSON(planets, &opts, writer)
		// Check if the content of the json is okay.
		var bytes []byte
		if bytes, err = ioutil.ReadFile(fileToWrite); err != nil {
			t.Fail()
		} else {
			allInOne := JSONReport{}
			if err := json.Unmarshal(bytes, &allInOne); err != nil {
				t.Fail()
			} else {
				if allInOne.Planets[0].ID != planet.id ||
					allInOne.Planets[0].Output != planet.outputStruct.output ||
					allInOne.Planets[0].Errored != planet.outputStruct.errored {
					t.Fail()
				}
			}
		}
	}

	// Tear down
	defer func() {
		os.Setenv("ORBIT_HOME", backup)
		os.RemoveAll(outDir)
	}()
}

// func createJSONReport(options map[string]string, planets []Planet, opts *Opts)
func TestCreateJSONReport(t *testing.T) {
	// Setup
	jobFile := "job"
	backup := os.Getenv("ORBIT_HOME")
	os.Setenv("ORBIT_HOME", os.TempDir())
	setupDirs()

	options := map[string]string{}
	options["job_name"] = path.Base(jobFile)
	options["orbit_home"] = os.Getenv("ORBIT_HOME")
	options["output"] = "reports"

	planet := Planet{
		id:         "app",
		name:       "App-Package 1",
		user:       "none",
		host:       "localhost",
		planetType: "server",
		dbID:       "",
		valid:      true,
		outputStruct: &StructuredOuput{
			planet:   "app",
			output:   "exit status 2",
			table:    nil,
			position: 0,
			errored:  true,
		},
	}
	planets := []Planet{planet}

	// "-j=/tmp/job.js", "-t=\"perlver_template\"", "-p", "-d=true", "app"
	opts := Opts{
		Template:  "perlver_template",
		Pretty:    false,
		Debug:     false,
		MaxToKeep: 2,
		Planets:   []string{"app"},
	}

	createJSONReport(options, planets, &opts)
	// Check if the content of the json is okay.
	jobName := strings.Split(options["job_name"], ".")[0]
	// ${ORBIT_HOME/cron_jobs/job/${timestamp_in_iso8601 with : replaced with _ }.json
	outputFolder := path.Join(options["orbit_home"], options["output"], jobName)
	latestReport := findLatest(outputFolder)
	fmt.Printf("Attempting to unmarshal JSONReport from %s\n", latestReport)
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(latestReport); err == nil {
		report := JSONReport{}
		if err = json.Unmarshal(bytes, &report); err == nil {
			wrapper := report.Planets[0]
			if wrapper.ID != planet.id ||
				wrapper.Output != planet.outputStruct.output ||
				wrapper.Errored != planet.outputStruct.errored {
				fmt.Fprintln(os.Stderr, "Unmarshalled object contains wrong values.")
				t.Fail()
			}
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		t.Fail()
	}

	// Tear down
	defer func() {
		os.Setenv("ORBIT_HOME", backup)
	}()
}

// Just a test function
func findLatest(outputFolder string) string {
	now := time.Now()
	format := "%d-%02d-%02dT%02d_%02d_%02d"
	stamp := fmt.Sprintf(format, now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fileToWrite := strings.Join([]string{stamp, "json"}, ".")
	return path.Join(outputFolder, fileToWrite)
}
