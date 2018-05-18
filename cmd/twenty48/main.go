package main

import (
	"flag"
)

var (
	fVersion bool

	fWidth  int
	fHeight int
)

func init() {
	flag.BoolVar(&fVersion, "version", false, "")
	flag.IntVar(&fWidth, "width", 8, "board width")
	flag.IntVar(&fHeight, "height", 8, "board height")
}

func main() {
	flag.Parse()

	if fVersion {
		PrintVersion()
		return
	}

	startGame(fWidth, fHeight)
}
