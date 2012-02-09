package main

import "gocurse"
import "fmt"

func main() {
	window, err := gocurse.Initscr()
	if err != nil {
		gocurse.Endwin()
		return
	}
	window.Refresh()
	fmt.Println(window.Getch())
	
	gocurse.Endwin()
}
