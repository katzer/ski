package main

import (
	"encoding/json"
	"fmt"
)

func decode(jsonObject string) [][]string {
	var jsonBlob = []byte(jsonObject)
	var toReturn = make([][]string, 0)
	err := json.Unmarshal(jsonBlob, &toReturn)
	if err != nil {
		fmt.Println("error:", err)
	}
	return toReturn
}
