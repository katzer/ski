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

func getFullConnectionDetails(skiString string) (string, string, string) {
	var dbID string
	split := strings.Split(skiString, "|")
	connDet := split[len(split)-1]
	user := strings.Split(connDet, "@")[0]
	host := strings.TrimSuffix(strings.Split(connDet, "@")[1], "\n")

	log.Debugf("skiString: %s", skiString)
	log.Debugf("user: %s", user)
	log.Debugf("host: %s", host)
	log.Debugf("dbID: %s", dbID)
	return user, host, dbID
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
*	checks, wether a planet is supported by ski or not
 */
func isSupported(planetType string) bool {
	switch planetType {
	case linuxServer:
		return true
	case database:
		return true
	case webServer:
		os.Stderr.WriteString("Usage of ski with web servers is not implemented")
		log.Fatalf("Usage of ski with web servers is not implemented")
		return false
	default:
		os.Stderr.WriteString("Unkown Type of target")
		log.Fatalf("Unkown Type of target")
		return false

	}
}

/**
*	Returns the type of a given planet
*	@params:
*		id: The planets id
*	@return: The planets type
 */
func getType(skiString string) string {

	return strings.Split(skiString, "|")[0]
}

/**
*	Returns the connection details to a given planet
*	@params:
*		id: The planets id
*	@return: The connection details to the planet
 */
func getFullSkiString(id string) string {
	cmd := exec.Command("fifa", "-f=ski", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		message := fmt.Sprintf("%s output is: %s called from ErrOut.\n", err, out)
		os.Stderr.WriteString("Unkown target")
		log.Fatalf(message)
	}
	return string(out)

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
