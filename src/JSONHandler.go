package main

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

func decode(planet *Planet, jsonObject string) ([][]string, error) {
	var jsonBlob = []byte(jsonObject)
	var toReturn = make([][]string, 0)
	err := json.Unmarshal(jsonBlob, &toReturn)
	if err != nil {
		log.Errorln(err)
		return make([][]string, 0), err
	}
	return toReturn, nil
}
