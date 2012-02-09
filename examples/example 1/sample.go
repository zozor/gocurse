package main

import "gocurse"
import "fmt"

func main() {
	//setup
	window, err := gocurse.Initscr()
	if err != nil {
		fmt.Println(err)
		gocurse.Endwin()
		return
	}
	defer gocurse.Endwin()
	
	window.Clear()
	
	gocurse.Noecho()
	gocurse.Cbreak()
	
	window.Keypad(true)
	
	if err := gocurse.StartColor(); err != nil {
		fmt.Println(err)
		return
	}
	
	//color setup
	gocurse.SetColorPair(1, gocurse.COLOR_RED, gocurse.COLOR_BLACK);
	
	x,y := 10, 10

	for {
		
		inp := window.Getch()
		fmt.Println(inp)

		switch inp {
		case 'q': return
		case gocurse.KEY_LEFT: x--
		case gocurse.KEY_RIGHT: x++
		case gocurse.KEY_UP: y--
		case gocurse.KEY_DOWN:y++
		}

		window.Clear()
		window.Addch(x, y, '@', gocurse.GetColorPair(1))
		window.Refresh()
	}
}
