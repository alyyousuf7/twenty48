package main

import (
	"fmt"
)

// build variables
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func PrintVersion() {
	fmt.Printf("version: %s\ncommit: %s\nbuild at: %s\n", version, commit, date)
}
