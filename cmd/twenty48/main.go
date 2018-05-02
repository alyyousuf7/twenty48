package main

import (
	"flag"
)

var (
	fWidth  int
	fHeight int
)

func init() {
	flag.IntVar(&fWidth, "width", 8, "board width")
	flag.IntVar(&fHeight, "height", 8, "board height")
}

func main() {
	flag.Parse()

	startGame(fWidth, fHeight)
}
