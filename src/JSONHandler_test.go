package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
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
	name := "cron_jobs"
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
				if allInOne.Planets[0].ID != planet.id {
					t.Fail()
				}
				if allInOne.Planets[0].Name != planet.name {
					t.Fail()
				}
				if allInOne.Planets[0].User != planet.user {
					t.Fail()
				}
				if allInOne.Planets[0].Host != planet.host {
					t.Fail()
				}
				if allInOne.Planets[0].PlanetType != planet.planetType {
					t.Fail()
				}
				if allInOne.Planets[0].DbID != planet.dbID {
					t.Fail()
				}
				if allInOne.Planets[0].Valid != planet.valid {
					t.Fail()
				}
				if allInOne.Planets[0].Output != planet.outputStruct.output {
					t.Fail()
				}
				if allInOne.Planets[0].Index != planet.outputStruct.position {
					t.Fail()
				}
				if allInOne.Planets[0].Errored != planet.outputStruct.errored {
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
