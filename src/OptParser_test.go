package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func setupOptParserTest() (string, error) {
	toUnMarshall := "job.json"
	// "-s=\"showver.sh\"", "-t=\"useless_template\"", "-d", "-p", "app"
	json := `{
  "debug"      : true,
  "help"       : false,
  "load"       : false,
  "version"    : false,
  "command"    : "",
  "pretty"     : true,
  "scriptName" : "\"showver.sh\"",
  "template"   : "\"useless_template\"",
  "planets"    : [ "app"]
}
`
	data := []byte(json)
	os.Mkdir(path.Join(os.TempDir(), "jobs"), 0744)
	absPath := path.Join(os.TempDir(), "jobs", toUnMarshall)

	if err := ioutil.WriteFile(absPath, data, 0644); err != nil {
		return "", err
	}
	return absPath, nil
}

func TestCreateTaskFromJobFile(t *testing.T) {
	fmt.Println("starting Test for CreateTaskFromJobFile")
	var absPath string
	var err error

	backup := os.Getenv("ORBIT_HOME")
	os.Setenv("ORBIT_HOME", os.TempDir())

	if absPath, err = setupOptParserTest(); err != nil {
		t.Fail()
	}

	fmt.Printf("reading file %s...\n", absPath)
	opts := createATaskFromJobFile(absPath)
	if opts.Help || opts.Version {
		fmt.Fprintf(os.Stderr, "parsed from json: %v", opts)
		t.Fail()
	}

	fmt.Println("ending Test for CreateTaskFromJobFile")

	defer func() {
		os.Setenv("ORBIT_HOME", backup)
		tearDownOptParserTest(absPath)
	}()
}

func tearDownOptParserTest(absPath string) {
	if err := os.Remove(absPath); err != nil {
		fmt.Fprintf(os.Stderr, "Could not remove the file %s after a test", absPath)
	}
}
