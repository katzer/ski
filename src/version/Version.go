package main

import (
	"os"
)

const version string = "0.9.2dev"

func main() {
	os.Stdout.WriteString(version)
}
