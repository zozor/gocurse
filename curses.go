package gocurse

// #cgo LDFLAGS: -lncurses
// struct ldat{};
// struct _win_st{};
// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
import "C"

import (
	"fmt"
	"unsafe"
	"errors"
)

var (
	ErrorInit = errors.New("Cannot initialize")
	ErrorWindowCreation = errors.New("Failed to create window")
	ErrorTerminalColor = errors.New("terminal does not support color")
	ErrorInitPair = errors.New("initializing color pair failed")
	ErrorNoEcho = errors.New("noecho failed")
	ErrorEcho = errors.New("echo failed")
	ErrorSetCursor = errors.New("curs_set failed")
	ErrorNocbreak = errors.New("Nocbreak failed")
	ErrorCbreak = errors.New("Cbreak failed")
	ErrorEndwin = errors.New("Endwin failed")
	ErrorKeypad = errors.New("Keypad failed")
	ErrorWinRefresh = errors.New("Window refresh failed")
	ErrorWinDelete = errors.New("window deletion failed")
	ErrorWindowMove = errors.New("window movement failed")
	ErrorWindowSyncok = errors.New("window syncok failed")
	ErrorNoInfo = errors.New("Something went wrong, ncurses does not know :(")
)

// Cursor options.
const (
	CURS_HIDE = iota;
	CURS_NORM;
	CURS_HIGH;
)

// Pointers to the values in curses, which may change values.
var Cols *int = nil;
var Rows *int = nil;

var Colors *int = nil;
var ColorPairs *int = nil;

var Tabsize *int = nil;

// Initializes gocurses
func init() {
	Cols = (*int)(unsafe.Pointer(&C.COLS));
	Rows = (*int)(unsafe.Pointer(&C.LINES));
	
	Colors = (*int)(unsafe.Pointer(&C.COLORS));
	ColorPairs = (*int)(unsafe.Pointer(&C.COLOR_PAIRS));
	
	Tabsize = (*int)(unsafe.Pointer(&C.TABSIZE));
}

/*
====================================================
Color Setup Functions
====================================================
*/
func StartColor() error {
	if C.has_colors() == 0 {
		return ErrorTerminalColor
	}
	C.start_color();
	
	return nil;
}

func SetColorPair(pair, fg, bg int) error {
	if C.init_pair(C.short(pair), C.short(fg), C.short(bg)) == 0 {
		return ErrorInitPair
	}
	return nil
}

func GetColorPair(pair int) int32 {
	return int32(C.COLOR_PAIR(C.int(pair)));
}

/*
====================================================
Curses Setup Functions
====================================================
*/

func Noecho() error {
	if int(C.noecho()) == 0 {
		return ErrorNoEcho
	}
	return nil;
}

func Echo() error {
	if int(C.echo()) == 0 {
		return ErrorEcho
	}
	return nil;
}

func SetCursor(c int) error {
	if C.curs_set(C.int(c)) == ERR {
		return ErrorSetCursor
	}
	return nil
}


func Nocbreak() error {
	if C.nocbreak() == ERR {
		return ErrorNocbreak
	}
	return nil;
}


func Cbreak() error {
	if C.cbreak() == ERR {
		return ErrorCbreak
	}
	return nil;
}

func Getch() int {
	return int(C.getch())
}


func Endwin() error {
	if C.endwin() == ERR {
		return ErrorEndwin
	}
	return nil;
}

func Raw() error {
	if C.raw() == ERR {
		return ErrorNoInfo //Forkert ERROR
	}
	return nil;
}

func Halfdelay(delay int) error {
	if C.halfdelay(C.int(delay)) == ERR {
		return ErrorNoInfo //Forkert ERROR
	}
	return nil
}

/*
====================================================
Type Window
====================================================
*/

type Window struct {
	win *C.WINDOW
}

func Initscr() (*Window, error) {
	win := C.initscr()

	if win == nil {
		return nil, ErrorInit
	}
	
	return &Window{win}, nil
}

func NewWindow(rows, cols, starty, startx int) (*Window, error) {
	nw := C.newwin(C.int(rows), C.int(cols), C.int(starty), C.int(startx))

	if nw == nil {
		return nil, ErrorWindowCreation
	}

	return &Window{nw}, nil;
}

func (w *Window) Delwindow() error {
	if C.delwin(w.win) == ERR {
		return ErrorWinDelete
	}
	return nil
}

func (w *Window) Getch() int {
	return int(C.wgetch(w.win))
}

func (w *Window) Getstr() (string, error) {
	var buffer [1024]C.char
	if C.wgetnstr(w.win, &buffer[0], 1024) == ERR {
		return "", ErrorNoInfo
	}
	
	s := C.GoString(&buffer[0])
	return s, nil
}

func (w *Window) Timeout(delay int) {
	C.wtimeout(w.win, C.int(delay))
}

func (w *Window) Nodelay(nodelay bool) {
	if nodelay {
		C.nodelay(w.win, 1)
	} else {
		C.nodelay(w.win, 0)
	}
}

func (w *Window) Addch(x, y int, c int32, flags int32) {
	C.mvwaddch(w.win, C.int(y), C.int(x), C.chtype(c) | C.chtype(flags));
}

// Since CGO currently can't handle varg C functions we'll mimic the
// ncurses addstr functions.
func (w *Window) Addstr(x, y int, str string, flags int32, v ...interface{}) {
	newstr := fmt.Sprintf(str, v...);
	
	w.MoveCursor(x, y);
	
	for i := 0; i < len(newstr); i++ {
		C.waddch(w.win, C.chtype(newstr[i]) | C.chtype(flags));
	}
}

func (w *Window) Border(ls, rs, ts, bs, tl, tr, bl, br int) error {
	if C.wborder(w.win, C.chtype(ls), C.chtype(rs),
	 C.chtype(ts), C.chtype(bs), C.chtype(tl), C.chtype(tr),
          C.chtype(bl), C.chtype(br)) == ERR {
          
          return ErrorNoInfo
    }
    return nil
}

// Normally Y is the first parameter passed in curses.
func (w *Window) MoveCursor(x, y int) error {
	if C.wmove(w.win, C.int(y), C.int(x)) == ERR {
		return ErrorWindowMove
	}
	return nil
}

func DublicateWindow(w *Window) *Window {
	return &Window{C.dupwin(w.win)}
}

//Flakey function according to ncurses DOC
func (w *Window) MoveDerwindow(x,y int) error {
	if C.mvderwin(w.win, C.int(y), C.int(x)) == ERR {
		return ErrorWindowMove
	}
	return nil
}



//Flakey function according to ncurses DOC
func (w *Window) Wsyncup() {
	C.wsyncup(w.win)
}

//Flakey function according to ncurses DOC
func (w *Window) Syncok(bf bool) error {
	var bfint int
	if bf {
		bfint++
	}
	
	if C.syncok(w.win, C.int(bfint)) == ERR {
		return ErrorWindowSyncok
	}
	return nil
}

//Flakey function according to ncurses DOC
func (w *Window) Wcursyncup() {
	C.wcursyncup(w.win)
}

//Flakey function according to ncurses DOC
func (w *Window) Wsyncdown() {
	C.wsyncdown(w.win)
}

//Flakey function according to ncurses DOC
func (w *Window) SubWindow(rows, cols, starty, startx int) (*Window, error)  {
	sw := C.subwin(w.win, C.int(rows), C.int(cols), C.int(starty), C.int(startx))

	if sw == nil {
		return nil, ErrorWindowCreation
	}

	return &Window{sw}, nil
}

//Flakey function according to ncurses DOC
func (w *Window) Derwin(rows, cols, starty, startx int) (*Window, error)  {
	dw := C.derwin(w.win, C.int(rows), C.int(cols), C.int(starty), C.int(startx))

	if dw == nil {
		return nil, ErrorWindowCreation
	}

	return &Window{dw}, nil;
}

func (w *Window) Keypad(tf bool) error {
	var outint int;
	if tf {
		outint++
	}
	if C.keypad(w.win, C.int(outint)) != OK {
		return ErrorKeypad
	}
	return nil
}

func (w *Window) Refresh() error {
	if C.wrefresh(w.win) == ERR {
		return ErrorWinRefresh
	}
	return nil
}

func (w *Window) Redrawln(beg_line, num_lines int) {
	C.wredrawln(w.win, C.int(beg_line), C.int(num_lines));
}

func (w *Window) Redraw() {
	C.redrawwin(w.win);
}

func (w *Window) Clear() {
	C.wclear(w.win);
}

func (w *Window) Erase() {
	C.werase(w.win);
}

func (w *Window) Clrtobot() {
	C.wclrtobot(w.win);
}

func (w *Window) Clrtoeol() {
	C.wclrtoeol(w.win);
}

func (w *Window) Box(verch, horch int) {
	C.box(w.win, C.chtype(verch), C.chtype(horch));
}

func (w *Window) Background(colour int32) {
	C.wbkgd(w.win, C.chtype(colour));
}

//Attributes
func (w *Window) Attroff(attrs int) error {
	if C.wattroff(w.win, C.int(attrs)) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Attron(attrs int) error {
	if C.wattron(w.win, C.int(attrs)) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Attrset(attrs int) error {
	if C.wattrset(w.win, C.int(attrs)) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Color_set(pair int) error {
	if C.wcolor_set(w.win, C.short(pair), nil) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Standend() error {
	if C.wstandend(w.win) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Standout() error {
	if C.wstandout(w.win) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) Attr_get() (int, int, error) {
	var attr C.attr_t
	var pair C.short
	
	if C.wattr_get(w.win, &attr, &pair, nil) == ERR {
		return 0,0, ErrorNoInfo
	}
	
	return int(attr), int(pair), nil
}

func (w *Window) Chgat(n, attr, color int) error {
	if C.wchgat(w.win, C.int(n), C.attr_t(attr), C.short(color), nil) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (w *Window) MoveChgat(x, y, n, attr, color int) error {
	if C.mvwchgat(w.win, C.int(y), C.int(x), C.int(n), C.attr_t(attr), C.short(color), nil) == ERR {
		return ErrorNoInfo
	}
	return nil
}

//Window sizes - Undefined for some reason?
/*
func (w *Window) Getyx() (x int, y int) {
	C.getyx(w.win, &y, &x)
	return
}

func (w *Window) Getparyx() (x int, y int) {
	C.getparyx(w.win, &y, &x)
	return
}

func (w *Window) Getbegyx() (x int, y int) {
	C.getbegyx(w.win, &y, &x)
	return
}

func (w *Window) Getmaxyx() (x int, y int) {
	C.getmaxyx(w.win, &y, &x)
	return
}
*/
