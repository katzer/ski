package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// Opts ...
type Opts struct {
	debugFlag      bool
	helpFlag       bool
	loadFlag       bool
	prettyFlag     bool
	scriptFlag     bool
	tableFlag      bool
	typeFlag       bool
	versionFlag    bool
	planetsCount   int
	bashScriptName string
	bashScriptPath string
	command        string
	currentDBDet   string
	currentDet     string
	dbScriptName   string
	dbScriptPath   string
	pyScriptName   string
	pyScriptPath   string
	scriptName     string
	scriptPath     string
	templateName   string
	templatePath   string
	planets        []string
}

const templatePath = `templates`
const dbScriptPath = `dbScripts`
const bashScriptPath = `bashScripts`
const pyScriptPath = `pythonScripts`

/**
*	Returns the contents of args in following order:
*	prettyprint flag
*	script flag
*	script path
*	command
*	planets
 */
func (opts *Opts) procArgs(args []string) {
	currentDir, err := os.Getwd()
	rootDir := path.Dir(currentDir)
	if err != nil {
		throwErr(err)
	}
	flag.BoolVar(&opts.helpFlag, "h", false, "help")
	flag.BoolVar(&opts.prettyFlag, "pp", false, "prettyprint")
	flag.BoolVar(&opts.typeFlag, "t", false, "type")
	flag.BoolVar(&opts.debugFlag, "d", false, "verbose")
	flag.BoolVar(&opts.loadFlag, "l", false, "ssh profile loading")
	flag.BoolVar(&opts.versionFlag, "v", false, "version")
	flag.StringVar(&opts.templatePath, "tp", path.Join(rootDir, templatePath), "path to template directory")
	flag.StringVar(&opts.templateName, "tn", "", "filename of template")
	flag.StringVar(&opts.bashScriptPath, "bp", path.Join(rootDir, bashScriptPath), "path to bash-script directory")
	flag.StringVar(&opts.dbScriptPath, "dp", path.Join(rootDir, dbScriptPath), "path to db-Script directory")
	flag.StringVar(&opts.pyScriptPath, "pyp", path.Join(rootDir, pyScriptPath), "path to python-Script directory")
	flag.StringVar(&opts.scriptName, "sn", "", "name of the script(regardless wether db or bash) to be executed")
	flag.StringVar(&opts.command, "c", "", "command to be executed in quotes")
	flag.StringVar(&opts.scriptPath, "s", "", "path to script file to be uploaded and executed")
	flag.Parse()
	opts.scriptFlag = !(opts.scriptPath == "")
	opts.tableFlag = !(opts.templateName == "")

	planets := flag.Args()
	opts.command = strings.TrimSuffix(strings.TrimPrefix(opts.command, "\""), "\"")
	opts.templateName = strings.TrimSuffix(strings.TrimPrefix(opts.templateName, "\""), "\"")
	opts.templatePath = strings.TrimSuffix(strings.TrimPrefix(opts.templatePath, "\""), "\"")
	opts.scriptPath = strings.TrimSuffix(strings.TrimPrefix(opts.scriptPath, "\""), "\"")
	opts.scriptName = strings.TrimSuffix(strings.TrimPrefix(opts.scriptName, "\""), "\"")
	opts.pyScriptPath = strings.TrimSuffix(strings.TrimPrefix(opts.pyScriptPath, "\""), "\"")
	for _, argument := range planets {
		if isSupported(argument) {
			opts.planets = append(opts.planets, argument)
			opts.planetsCount++
		} else {
			fmt.Fprintf(os.Stderr, "This Type of Connection is not supported.")
			os.Exit(1)
		}
	}
	if len(args) == 1 {
		opts.helpFlag = true
	}

}

/**
*	Splits the given connectiondetails and returns the hostname
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: hostname
 */
func getHost(connDet string) string {
	toReturn := strings.Split(connDet, "@")
	return toReturn[1]
}

/**
*	Splits the given connectiondetails and returns the user
*	@params:
*		connDet: Connection details in following form: user@hostname
*	@return: user
 */
func getUser(connDet string) string {
	toReturn := strings.Split(connDet, "@")
	return toReturn[0]
}

/**
*	Returns the type of a given planet
*	@params:
*		id: The planets id
*	@return: The planets type
 */
func getType(id string) string {
	cmd := exec.Command("ff", "-t", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		throwErrOut(out, err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	Returns the connection details to a given planet
*	@params:
*		id: The planets id
*	@return: The connection details to the planet
 */
func getPlanetDetails(id string) string {
	cmd := exec.Command("ff", id, "-f=pqdb")
	out, err := cmd.CombinedOutput()
	if err != nil {
		throwErrOut(out, err)
	}
	return strings.TrimSpace(string(out))
}

/**
*	counts the supported planets in a list of planets
 */
func countSupported(planets []string) int {
	i := 0
	for _, planet := range planets {
		if getType(planet) == "server" {
			i++
		}
	}
	return i
}

/**
*	checks, wether a planet is supported by goo or not
 */
func isSupported(planet string) bool {
	switch getType(planet) {
	case "server":
		return true
	case "db":
		return true
	case "web":
		return false
	default:
		return false

	}

}

func getMaxLength(toProcess string) int {
	_ = toProcess
	return 0
}

/**
*	splits db details (dbID:user@host) and returns them as dbID,user@host
 */
func procDBDets(dbDet string) (string, string) {
	parts := strings.Split(dbDet, ":")
	return parts[0], parts[1]
}
