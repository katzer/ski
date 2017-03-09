package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// codebeat:disable[TOO_MANY_IVARS]

// Opts structure for holding commandline arguments
type Opts struct {
	Debug      bool     `json:"debug"`
	Help       bool     `json:"help"`
	Load       bool     `json:"load"`
	Pretty     bool     `json:"pretty"`
	Version    bool     `json:"version"`
	SaveReport bool     `json:"save_report"`
	Command    string   `json:"command"`
	ScriptName string   `json:"scriptName"`
	Template   string   `json:"template"`
	Planets    []string `json:"planets"`
	LogFile    string   `json:"log_file"`
}

// codebeat:enable[TOO_MANY_IVARS]

func (opts *Opts) String() string {
	template := `opts : {
	Debug: %t
	Help: %t
	Load: %t
	Pretty: %t
	Version: %t
	SaveReport: %t
	Command: %s
	ScriptName: %s
	Template : %s
	Planets: %v
}
`

	return fmt.Sprintf(template,
		opts.Debug,
		opts.Help,
		opts.Load,
		opts.Pretty,
		opts.Version,
		opts.SaveReport,
		opts.Command,
		opts.ScriptName,
		opts.Template,
		opts.Planets)
}

func (opts *Opts) postProcessing() {
	// TODO ask what is happening here
	opts.Command = strings.TrimSuffix(strings.TrimPrefix(opts.Command, "\""), "\"")
	opts.Template = strings.TrimSuffix(strings.TrimPrefix(opts.Template, "\""), "\"")
	opts.ScriptName = strings.TrimSuffix(strings.TrimPrefix(opts.ScriptName, "\""), "\"")
}

func (opts *Opts) validate() {
	opts.validateExtension()
	opts.validateCommandAndScript()
	opts.checkForInvalidIds()
	opts.validateTemplate()
}

func (opts *Opts) checkForInvalidIds() {
	for _, id := range opts.Planets {
		// Check if any flags were given after planet ids, if yes stop the app
		if strings.HasPrefix(id, "-") {
			fmt.Fprintf(os.Stderr, "Unknown target: %s", id)
			os.Exit(1)
		}
	}
}

func (opts *Opts) validateCommandAndScript() {
	scriptNotEmpty := len(opts.ScriptName) > 0
	cmdNotEmpty := len(opts.Command) > 0
	if scriptNotEmpty && cmdNotEmpty {
		message := "providing both a script AND a command is not possible"
		err := errors.New(message)
		fmt.Fprintf(os.Stderr, "%s\nAddInf: %s\n", err, message)
		log.Fatal(err)
	}
}

//Checks if the given script got an acceptable extension
func (opts *Opts) validateExtension() {
	script := opts.ScriptName
	isSH := strings.HasSuffix(script, ".sh")
	isSQL := strings.HasSuffix(script, ".sql")
	if script != "" && !(isSH || isSQL) {
		message := fmt.Sprintf("The provided scripts %s must have either .sh or .sql extension.", script)
		fmt.Fprintln(os.Stderr, message)
		log.Fatal(message)
	}
}

func (opts *Opts) validateTemplate() {
	template := opts.Template != ""
	file := path.Join(os.Getenv("ORBIT_HOME"), "templates", opts.Template)
	if template {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			message := fmt.Sprintf("The provided template does not exist")
			fmt.Fprintln(os.Stderr, message)
			log.Fatal(message)
		}
	}

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
	bytes, err := ioutil.ReadFile(wcopy)
	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't open job file: %s", jsonFile)
		fmt.Fprint(os.Stderr, errorMessage)
		log.Fatal(errorMessage)
	}

	err = json.Unmarshal(bytes, &job)
	if err != nil {
		errorMessage := fmt.Sprintf("Error parsing job json file: %s", jsonFile)
		fmt.Fprint(os.Stderr, errorMessage)
		log.Fatal(errorMessage)
	}

	log.Debugf("Read a task from %s", jsonFile)
	log.Debugln()
	log.Debugf("Unmarshalled %v", job)
	log.Debugln()
	return job
}
