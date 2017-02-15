package main

import (
	"os"
)

const version string = "0.9.1"

func main() {
	os.Stdout.WriteString(version)
}
