package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"time"
	"strings"
	"math/rand"
)

var score = 0

var colon string = `
  §
##§
  §
##§
  §
`

var upArrow string = `
       §
   -   §
  /@\  §
 /@@@\ §
/@@@@@\§
       §
       §
`

var downArrow string = `
       §
       §
\@@@@@/§
 \@@@/ §
  \@/  §
   -   §
       §
`

var leftArrow string = `
       §
    /@ §
  /@@@ §
<@@@@@ §
  \@@@ §
    \@ §
       §
`

var rightArrow string = `
       §
@\     §
@@@\   §
<@@@@> §
@@@/   §
@/     §
       §
`

var zero string = `
###### §
#    # §
#    # §
#    # §
###### §
`
var one string = `
     # §
     # §
     # §
     # §
     # §
`
var two string = `
###### §
     # §
###### §
#      §
###### §
`

var three string = `
###### §
     # §
   ### §
     # §
###### §
`

var four string = `
#      §
#      §
#   #  §
###### §
    #  §
`

var five string = `
###### §
#      §
###### §
     # §
###### §
`

var six string = `
###### §
#      §
###### §
#    # §
###### §
`

var seven string = `
###### §
     # §
     # §
     # §
     # §
`

var height string = `
###### §
#    # §
###### §
#    # §
###### §
`

var nine string = `
###### §
#    # §
###### §
     # §
###### §
`

func printSymbol(x, y int, fg, bg termbox.Attribute, msg string) {
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

func printBorders() {

	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if x == 0 || x == w-1 {
				printSymbol(x, y, termbox.ColorWhite, termbox.ColorDefault, "#")
			} else {
				if y > 0 && y < h-1 {
					printSymbol(x, y, termbox.ColorWhite, termbox.ColorDefault, " ")
				} else {
					printSymbol(x, y, termbox.ColorWhite, termbox.ColorDefault, "#")
				}
			}
		}
	}

}


func printString(s string, c termbox.Attribute) {
	printBorders()

	//termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	w, h := termbox.Size()
	s = fmt.Sprintf("%s", s)
	ds := strings.Split(s, "§")
	d := -1
	for _, s := range ds {
		if len(s) > d {
			d = len(s)
		}
	}

	printSymbol((w/2)-(d/2), (h-strings.Count(s, "§"))/2, c, termbox.ColorDefault, s)
}

func timer() {

	time.Sleep(1 * time.Second)
	printString(three, termbox.ColorRed)
	time.Sleep(1 * time.Second)
	printString(two, termbox.ColorRed)
	time.Sleep(1 * time.Second)
	printString(one, termbox.ColorRed)
	time.Sleep(1 * time.Second)

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func arrows(arrow int) bool {
	switch arrow {
		case 0:
			printString(upArrow, termbox.ColorRed)
		case 1:
			printString(downArrow, termbox.ColorRed)
		case 2:
			printString(leftArrow, termbox.ColorRed)
		case 3:
			printString(rightArrow, termbox.ColorRed)
	}
	switch ev := termbox.PollEvent(); ev.Key {
		case termbox.KeyArrowUp:
			printString(upArrow, termbox.ColorGreen)
			time.Sleep(200 * time.Millisecond)
			return arrow == 0
		case termbox.KeyArrowDown:
			printString(downArrow, termbox.ColorGreen)
			time.Sleep(200 * time.Millisecond)
			return arrow == 1
		case termbox.KeyArrowLeft:
			printString(leftArrow, termbox.ColorGreen)
			time.Sleep(200 * time.Millisecond)
			return arrow == 2
		case termbox.KeyArrowRight:
			printString(rightArrow, termbox.ColorGreen)
			time.Sleep(200 * time.Millisecond)
			return arrow == 3
		default :
			return false
	}
	return false
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	printString(
`                       WELCOME TO....                      §` +
`                          ARROWS!!!                           §` +
`                                                              §` +
`    There is only 1 rule: press what appears on the screen!   §` +
`    Suggestion: be as fast as possible and DON'T press random §` +
`                                                              §` +
`    Commands:                                                 §` +
`         ESC:  anytime to exit game;                          §` +
`         Char: any char that appears on the screen            §` +
`                                                              §` +
`    Press Ctrl+S when you are ready to start...               §`, termbox.ColorYellow)

	termbox.SetInputMode(termbox.InputEsc)

	mainloop:
		for {
			//myrand := random(1, 9)
			switch ev := termbox.PollEvent(); ev.Key {
			case termbox.KeyCtrlS:
				timer()
				arrowloop:
				for {
					goon := arrows(random(0, 3))
					if !goon {
						printString(fmt.Sprintf(
							`_______  _______  __   __  _______    _______  __   __  _______  ______§
							|       ||   _   ||  |_|  ||       |  |       ||  | |  ||       ||    _ |  §
							|    ___||  |_|  ||       ||    ___|  |   _   ||  |_|  ||    ___||   | ||  §
							|   | __ |       ||       ||   |___   |  | |  ||       ||   |___ |   |_||_ §
							|   ||  ||       ||       ||    ___|  |  |_|  ||       ||    ___||    __  |§
							|   |_| ||   _   || ||_|| ||   |___   |       | |     | |   |___ |   |  | |§
							|_______||__| |__||_|   |_||_______|  |_______|  |___|  |_______||___|  |_|§
							|_______||__| |__||_|   |_||_______|  |_______|  |___|  |_______||___|  |_|§
                                                                                                       §
												         YOUR FINAL SCORE: %d
							`, score), termbox.ColorRed)

						time.Sleep(3 * time.Second)

						printString(
							`                       WELCOME TO....                      §` +
								`                          ARROWS!!!                           §` +
								`                                                              §` +
								`    There is only 1 rule: press what appears on the screen!   §` +
								`    Suggestion: be as fast as possible and DON'T press random §` +
								`                                                              §` +
								`    Commands:                                                 §` +
								`         ESC:  anytime to exit game;                          §` +
								`         Char: any char that appears on the screen            §` +
								`                                                              §` +
								`    Press Ctrl+S when you are ready to start...               §`, termbox.ColorYellow)
						break arrowloop
					} else {
						score++
					}
					fmt.Printf("SCORE: %d", score)
				}
			case termbox.KeyEsc:
				break mainloop
			}
		}
	termbox.Close()
	fmt.Println("Finished!")
}
