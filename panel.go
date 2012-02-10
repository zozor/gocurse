package gocurse

//#cgo LDFLAGS: -lpanel
//#define _Bool int
//#include <panel.h>
import "C"


func UpdatePanels() {
	C.update_panels()
}

type Panel struct {
	panel *C.PANEL
}

func (p *Panel) Window() *Window {
	return &Window{C.panel_window(p.panel)}
}

func (p *Panel) Hide() error {
	if C.hide_panel(p.panel) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Show() error {
	if C.show_panel(p.panel) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Del() error {
	if C.del_panel(p.panel) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Top() error {
	if C.top_panel(p.panel) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Bottom() error {
	if C.bottom_panel(p.panel) == ERR {
		return Error
	}
	return nil
}

func (w *Window) NewPanel() *Panel {
	return &Panel{C.new_panel(w.win)}
}

func (p *Panel) Above() *Panel {
	panel := C.panel_above(p.panel)
	if panel == nil {
		return nil
	}
	return &Panel{panel}
}

func (p *Panel) Below() *Panel {
	panel := C.panel_below(p.panel)
	if panel == nil {
		return nil
	}
	return &Panel{panel}
}

//extern NCURSES_EXPORT(int)     set_panel_userptr (PANEL *, NCURSES_CONST void *);
//extern NCURSES_EXPORT(NCURSES_CONST void*) panel_userptr (const PANEL *);

func (p *Panel) Move(x,y int) error {
	if C.move_panel(p.panel, C.int(y),C.int(x)) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Replace(win *Window) error {
	if C.replace_panel(p.panel, win.win) == ERR {
		return Error
	}
	return nil
}

func (p *Panel) Hidden() bool {
	if C.panel_hidden(p.panel) == 0 {
		return false
	}
	return true
}
