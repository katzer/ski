package main

import (
	"fmt"

	"os"
	"os/exec"
	"path"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

// StructuredOuput ...
type StructuredOuput struct {
	planet       string
	output       string
	maxOutLength int
}

func main() {
	opts := Opts{}
	opts.procArgs(os.Args)

	if opts.helpFlag {
		printUsage()
		os.Exit(0)
	}
	if opts.versionFlag {
		printVersion()
		os.Exit(0)
	}

	level := log.InfoLevel
	if opts.debugFlag {
		level = log.DebugLevel
	}
	// Default logfile path
	logDir := path.Join(os.Getenv("ORBIT_HOME"), "logs")
	createLogDirIfNecessary(logDir)
	logFile := path.Join(logDir, "ski.log")
	setupLogger(logFile, level)

	log.Infof("Started with args: %v", os.Args)
	log.Debug(&opts)
	exec := makeExecutor(&opts)
	exec.execMain(&opts)
	log.Infof("Ended with args: %v", os.Args)
}

func createLogDirIfNecessary(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0775|os.ModeDir); err != nil {
			// can't do anything
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
		}
	}
}

func makeExecutor(opts *Opts) Executor {
	log.Debugf("Function: makeExecutor")
	executor := Executor{}
	for _, entry := range opts.planets {
		fullSkiString := getFullSkiString(entry)
		var planet Planet
		planet.planetType = getType(fullSkiString)
		if !isSupported(planet.planetType) {
			continue
		}
		planet.id = entry
		planet.user, planet.host, planet.dbID = getFullConnectionDetails(fullSkiString)
		planet.outputStruct = StructuredOuput{entry, "", 0}
		executor.planets = append(executor.planets, planet)
	}
	log.Debugf("executor: %s", executor)
	return executor
}

/**
*	Prints the help dialog
 */
func printUsage() {
	usage := `usage: ski [options...] <planets>...
	Options:
	-s="<scriptname>"   Execute script and return result
	-c="<command>"      Execute script and return result
	-t=<"templatename>" Templatefile to be applied
	-p    Pretty print output as a table
	-l    Load bash profiles on Server
	-h    Display this help text
	-v    Show version number
	-d    Show extended debug informations, set logging level to debug
`
	fmt.Println(usage)
}

func getOS() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "darwin":
		return "MacOS"
	default:
		return "could not determine OS"
	}
}

func getArch() string {
	switch runtime.GOARCH {
	case "amd64":
		return "64bit"
	case "386":
		return "32bit"
	default:
		return "could not determine architecture"
	}
}

func getOSArch() string {
	switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("uname", "-m").Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		return string(out)
	case "windows":
		out, err := exec.Command("if exist \"%ProgramFiles(x86)%\" echo 64-bit").Output()
		if err != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err)
		}
		if string(out) == "64-bit" {
			return "x86_64"
		}
		return "i686"
	default:
		return "could not determine Operating system"
	}
}
