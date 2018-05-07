package main

import (
	"fmt"
	"os"

	"github.com/alyyousuf7/twenty48"
	termbox "github.com/nsf/termbox-go"
)

func startGame(width, height int) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	board := twenty48.NewBoard(width, height)

	for {
		board.NewMove()
		drawBoard(board, os.Stdout)

		switch ev := termbox.PollEvent(); {
		case ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowUp:
			board.MoveUp()
		case ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowDown:
			board.MoveDown()
		case ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowLeft:
			board.MoveLeft()
		case ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowRight:
			board.MoveRight()
		case ev.Type == termbox.EventKey && (ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC):
			return
		}
	}
}

func drawCell(x, y, cellWidth int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	if text == "0" {
		text = ""
	}
	for i := 0; i < cellWidth; i++ {
		char := rune(' ')
		if i < len(text) {
			char = rune(text[i])
		}
		termbox.SetCell(x+i, y, char, fg, bg)
	}
}

func drawBoard(board *twenty48.Board, file *os.File) {
	cellWidth := 4

	w, h := board.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := board.At(x, y).Value()

			fg := termbox.ColorWhite
			bg := termbox.ColorDefault
			switch v {
			case 2:
				bg = termbox.ColorBlack
			case 4:
				bg = termbox.ColorBlue
			case 8:
				bg = termbox.ColorCyan
			case 16:
				bg = termbox.ColorDefault
			case 32:
				bg = termbox.ColorGreen
			case 64:
				bg = termbox.ColorMagenta
			case 128:
				bg = termbox.ColorRed
			case 256:
				fg = termbox.ColorBlack
				bg = termbox.ColorWhite
			case 512:
				bg = termbox.ColorYellow
			case 1024:
				bg = termbox.ColorBlack
			case 2048:
				bg = termbox.ColorBlue
			}
			drawCell(x*cellWidth, y, cellWidth, fmt.Sprintf("%d", v), fg, bg)
		}
	}

	termbox.Flush()
}
