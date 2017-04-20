package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

// Test to see if the flags in job file override/ignore
// the ones given as cmdline args
func TestParseOptions(t *testing.T) {
	// create the json file in tmp folder
	jsonFile, err := setupOptParserTest()
	if err != nil {
		t.Fail()
	}

	backup := os.Getenv("ORBIT_HOME")
	os.Setenv("ORBIT_HOME", os.TempDir())
	// remove the json file from the tmp folder
	defer func() {
		os.Setenv("ORBIT_HOME", backup)
		tearDownOptParserTest(jsonFile)
	}()

	_, basename := path.Split(jsonFile)
	onlyName := strings.Split(basename, ".")[0]
	jobOption := fmt.Sprintf("-j=%s", onlyName)
	os.Args = []string{"ski", jobOption, "-d=false", "-v=false", "-c=\"ls -la ${HOME}\""}
	fmt.Printf("TestParseOptions :: os.Args: %v\n", os.Args)
	opts := parseOptions()

	fmt.Printf("TestParseOptions :: opts: %s\n", opts.String())

	errors := make([]string, 0)
	if !opts.Debug {
		errors = append(errors, "Debug flag was set through the cmdline option.")
	}
	if opts.Version {
		errors = append(errors, "Version flag set through the cmdline option.")
	}
	if opts.Command != "" {
		msg := fmt.Sprintf("Command was not parsed correctly: %s", opts.Command)
		errors = append(errors, msg)
	}

	if len(errors) > 0 {
		for i, message := range errors {
			fmt.Fprintf(os.Stdout, "%d. Error: %s\n", i, message)
		}
		t.Fail()
	}
}
