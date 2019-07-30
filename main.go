package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}

func writeString(x, y int, text string) {
	offset := x
	for _, r := range []rune(text) {
		termbox.SetCell(offset, y, r, termbox.ColorDefault, termbox.ColorDefault)
		offset += runewidth.RuneWidth(r)
	}

}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	writeString(1, 1, "Simple Counter")
	writeString(1, 2, "Return or Space: Count up, Ctrl-R: Reset, Esc: Quit")

	termbox.Flush()
}

func main() {

	err := termbox.Init()
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	defer termbox.Close()

	clear()

	count := 0
loop:
	for {
		writeString(1, 4, fmt.Sprintf("\rCount: %d", count))
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				fmt.Println("\nQuic.")
				break loop
			case termbox.KeyEnter, termbox.KeySpace:
				count++
			case termbox.KeyCtrlR:
				count = 0
				clear()
			}
		}
	}
}
