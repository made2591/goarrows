package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	lx := -1
	for _, c := range msg {
		if c == '§' {
			y++
			if lx != -1 {
				x = lx
			} else {
				lx = x
			}
		} else {
			termbox.SetCell(x, y, c, fg, bg)
			if lx == -1 {
				lx = x
			}
			x++
		}
	}
}

func draw(i int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()
	s := fmt.Sprintf("count = %d", i)

	tbPrint((w/2)-(len(s)/2), h/2, termbox.ColorRed, termbox.ColorDefault, s)
}

func drawS(s string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()
	s = fmt.Sprintf("%s", s)

	tbPrint((w/2)-(len(s)/2), h/2, termbox.ColorRed, termbox.ColorDefault, s)
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	go func() {
		time.Sleep(5 * time.Second)
		termbox.Interrupt()

		// This should never run - the Interrupt(), above, should cause the event
		// loop below to exit, which then exits the process.  If something goes
		// wrong, this panic will trigger and show what happened.
		time.Sleep(1 * time.Second)
		panic("this should never run")
	}()

	drawS("   /\\§  /__\\§   ||§   ||§")
	mainloop:
		for {
			switch ev := termbox.PollEvent(); ev.Key {
			case termbox.KeyArrowUp:
				drawS("   /\\§  /__\\§   ||§   ||§")
			case termbox.KeyArrowDown:
				drawS("   ||§   ||§  \\--/§   \\/ §")
			case termbox.KeyArrowLeft:
				drawS("⬅")
			case termbox.KeyArrowRight:
				drawS("->")
			case termbox.KeyEsc:
				break mainloop
			}
		}
		termbox.Close()

		fmt.Println("Finished")
}
