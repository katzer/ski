package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func parseConnectionDetails(ids []string) []Planet {
	// NOTE: fifa swapped type and id positions, id comes first
	skiStrings := getFullSkiString(ids)
	retval := make([]Planet, 0)
	for i, skiString := range skiStrings {
		tokens := strings.Split(skiString, skiDelim)
		connectionURL := tokens[len(tokens)-1]
		urlTokens := strings.Split(connectionURL, ":")

		var dbID string
		planetID, planetType, name := tokens[0], tokens[1], tokens[2]
		if len(urlTokens) > 1 {
			dbID = urlTokens[0]
		}

		user, host := getUserAndHost(connectionURL)

		planet := Planet{
			id:         planetID,
			planetType: planetType,
			name:       name,
			dbID:       dbID,
			user:       user,
			host:       host,
			outputStruct: &StructuredOuput{
				planet:   planetID,
				output:   "",
				table:    make([][]string, 0),
				position: i,
                errored : false,
			},
		}

		planet.valid = isValidPlanet(planet)
		// TODO Write the ski string to out?
		// if (!planet.valid) {
		// }
		log.Debugf("skiString: %s, and planet parsed from it: %v", skiString, planet)
		retval = append(retval, planet)
	}
	return retval
}

func getKeyPath() string {
	keyPath := os.Getenv("ORBIT_KEY")
	if keyPath == "" {
		if runtime.GOOS == windows {
			keyPath = ""
		} else {
			keyPath = path.Join(os.Getenv("ORBIT_HOME"), "config", "ssh", "orbit.key")
		}
	}
	return strings.TrimPrefix(keyPath, os.Getenv("HOME"))
}

func isSupported(planetType string) bool {
	supported := map[string]bool{database: true, linuxServer: true, webServer: false}
	return supported[planetType]
}

/**
*	Returns the connection details to a given planet
*	@params:
*		id: The planets id
*	@return: The connection details to the planet
 */
func getFullSkiString(ids []string) []string {
	length := len(ids)
	if length == 0 {
		return []string{}
	}

	args := append([]string{"-f=ski"}, ids...)
	cmd := exec.Command("fifa", args...)
	// TODO check the exit code etc. if len(cmd.Path) == 0 {}
	out, err := cmd.CombinedOutput()
	skiFormat := validateSkiFormat(string(out))
	if !skiFormat {
		message := "fifa output is not valid " + string(out)
		log.Fatalf(message)
	}
	if err != nil {
		message := fmt.Sprintf("%s output is: %s called from ErrOut.\n", err, out)
		log.Warnf(message)

	}
	// NOTE: "\n" at the end
	wcopy := strings.TrimSuffix(string(out), "\n")
	lines := strings.Split(wcopy, "\n")
	for i, line := range lines {
		log.Debugf("%d lines received.", length)
		log.Debugf("Line %d: %s\n", i, line)
	}
	return lines
}

func getUserAndHost(connectionURL string) (string, string) {
	var tokens []string
	idxCol := strings.IndexRune(connectionURL, ':')
	idxAt := strings.IndexRune(connectionURL, '@')
	if idxAt < 0 {
		log.Warnf("invalid address: %s", connectionURL)
		return "", ""
	}
	if idxCol < 0 {
		tokens = strings.Split(connectionURL, "@")
		return tokens[0], tokens[1]
	}
	tokens = strings.Split(connectionURL[idxCol+1:], "@")
	return tokens[0], tokens[1]
}

func validateSkiFormat(fifaString string) bool {
	firstLine := strings.Split(fifaString, "\n")[0]
	tokens := strings.Split(firstLine, skiDelim)
	return len(tokens) >= 4
}
