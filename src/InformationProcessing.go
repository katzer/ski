package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func getFullConnectionDetails(planet string) (string, string) {
	var dbID, connDet string
	planetType := getType(planet)
	connDet = getPlanetDetails(planet)
	switch planetType {
	case linuxServer:
		dbID = ""
	case database:
		dbID, connDet = procDBDets(connDet)
	case webServer:
		message := "Usage of ski with web servers is not implemented"
		os.Stderr.WriteString(message)
		log.Warnln(message)
	default:
		message := fmt.Sprintf("Unkown Type of target %s: %s\n", planet, planetType)
		os.Stderr.WriteString(message)
		log.Warnln(message)
	}
	return dbID, connDet
}

/**
*	splits db details (dbID:user@host) and returns them as dbID,user@host
 */
func procDBDets(dbDet string) (string, string) {
	parts := strings.Split(dbDet, ":")
	return parts[0], parts[1]
}

/**
*	Returns the proper Keypath
 */
func getKeyPath() string {
	keyPath := os.Getenv("ORBIT_KEY")
	if keyPath == "" {
		if runtime.GOOS == "windows" {
			keyPath = os.Getenv("TEMP") + "\\tempTabFormat.py"
		} else {
			keyPath = path.Join(os.Getenv("ORBIT_HOME"), "config", "ssh", "orbit.key")
		}
	}
	return strings.TrimPrefix(keyPath, os.Getenv("HOME"))
}

/**
*	counts the supported planets in a list of planets
 */
func countSupported(planets []string) int {
	i := 0
	for _, planet := range planets {
		if getType(planet) == linuxServer {
			i++
		}
	}
	return i
}

/**
*	checks, wether a planet is supported by ski or not
 */
func isSupported(planet string) bool {
	supported := map[string]bool{database: true, linuxServer: true, webServer: false}
	planetType := getType(planet)
	return supported[planetType]
}

/**
*	Returns the type of a given planet
*	@params:
*		id: The planets id
*	@return: The planets type
 */
func getType(id string) string {
	cmd := exec.Command("fifa", "-t", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		message := fmt.Sprintf("%s output is: %s called from ErrOut.\n", err, out)
		os.Stderr.WriteString(message)
		log.Fatalln(message)
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
	cmd := exec.Command("fifa", "-f=ski", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		message := fmt.Sprintf("%s output is: %s called from ErrOut.\n", err, out)
		os.Stderr.WriteString(message)
		log.Fatalln(message)
	}
	if len(strings.Split(string(out), "|")) >= 4 {
		return strings.TrimSpace(strings.Split(string(out), "|")[3])
	}
	return ""

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
func getUserAndHost(connDet string) (string, string) {
	// TODO: error handling or remove the func completely
	toReturn := strings.Split(connDet, "@")
	return toReturn[0], toReturn[1]
}
