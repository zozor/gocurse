package gocurse

//#cgo LDFLAGS: -lmenu
//#define _Bool int
//#include <menu.h>
import "C"

type Text struct {
	text *C.TEXT
}

type Item struct {
	item *C.ITEM
}

type Menu struct {
	menu *C.MENU
}

type ItemOptions struct {
	itemOptions C.Item_Options
}

type MenuOptions struct {
	menuOptions C.Menu_Options
}


const (
    O_ONEVALUE = C.O_ONEVALUE;
    O_SHOWDESC = C.O_SHOWDESC;
    O_ROWMAJOR = C.O_ROWMAJOR;
    O_IGNORECASE = C.O_IGNORECASE;
    O_SHOWMATCH = C.O_SHOWMATCH;
    O_NONCYCLIC = C.O_NONCYCLIC;
    O_SELECTABLE = C.O_SELECTABLE;
    REQ_LEFT_ITEM = C.REQ_LEFT_ITEM;
    REQ_RIGHT_ITEM = C.REQ_RIGHT_ITEM;
    REQ_UP_ITEM = C.REQ_UP_ITEM;
    REQ_DOWN_ITEM = C.REQ_DOWN_ITEM;
    REQ_SCR_ULINE = C.REQ_SCR_ULINE;
    REQ_SCR_DLINE = C.REQ_SCR_DLINE;
    REQ_SCR_DPAGE = C.REQ_SCR_DPAGE;
    REQ_SCR_UPAGE = C.REQ_SCR_UPAGE;
    REQ_FIRST_ITEM = C.REQ_FIRST_ITEM;
    REQ_LAST_ITEM = C.REQ_LAST_ITEM;
    REQ_NEXT_ITEM = C.REQ_NEXT_ITEM;
    REQ_PREV_ITEM = C.REQ_PREV_ITEM;
    REQ_TOGGLE_ITEM = C.REQ_TOGGLE_ITEM;
    REQ_CLEAR_PATTERN = C.REQ_CLEAR_PATTERN;
    REQ_BACK_PATTERN = C.REQ_BACK_PATTERN;
    REQ_NEXT_MATCH = C.REQ_NEXT_MATCH;
    REQ_PREV_MATCH = C.REQ_PREV_MATCH;

    MIN_MENU_COMMAND = C.MIN_MENU_COMMAND;
    MAX_MENU_COMMAND = C.MAX_MENU_COMMAND;
)

func (m *Menu) CurrentItem() *Item {
    return &Item{C.current_item(m.menu)}
}

func NewItem(name string, desc string) *Item { //Memory Leaks: do something about it sometime
    return &Item{C.new_item(C.CString(name), C.CString(desc))}
}

func NewMenu(items []*Item) (*Menu, error) {
	item := make([]*C.ITEM, 0, len(items))
	for _, i := range items {
		item = append(item, i.item)
	}
	
	menu := &Menu{C.new_menu(&item[0])}
	
	if menu == nil {
		return nil, ErrorNoInfo
	}
	return menu, nil
}

func (i *Item) Opts() ItemOptions {
    return ItemOptions{C.item_opts(i.item)}
}

func (m *Menu) Opts() MenuOptions {
    return MenuOptions{C.menu_opts(m.menu)}
}

func (i *Item) Description() string {
    return C.GoString(C.item_description(i.item))
}

func (i *Item) Name() string {
    return C.GoString(C.item_name(i.item))
}

func (m *Menu) Mark() string {
    return C.GoString(C.menu_mark(m.menu))
}

func (m *Menu) SetMark(mark string) error {
	if C.set_menu_mark(m.menu, C.CString(mark)) == ERR {
		return ErrorNoInfo
	}
    return nil
}

func (m *Menu) Pattern() string {
    return C.GoString(C.menu_pattern(m.menu))
}

func (m *Menu) Back() int {
	return int(C.menu_back(m.menu))
}

func (m *Menu) Fore() int {
	return int(C.menu_fore(m.menu))
}

func (m *Menu) Grey() int {
	return int(C.menu_grey(m.menu))
}

func (i *Item) Free() error {
	if C.free_item(i.item) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) Free() error {
	if C.free_menu(m.menu) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) ItemCount() int {
	return int(C.item_count(m.menu))
}

func (i *Item) Index() int {
	return int(C.item_index(i.item))
}

func (i *Item) OptsOn(opt ItemOptions) error {
	if C.item_opts_on(i.item, opt.itemOptions) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (i *Item) OptsOff(opt ItemOptions) error {
	if C.item_opts_off(i.item, opt.itemOptions) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) Drive(req int) error {
	if C.menu_driver(m.menu, C.int(req)) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) OptsOn(opt MenuOptions) error {
	if C.menu_opts_on(m.menu, opt.menuOptions) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) OptsOff(opt MenuOptions) error {
	if C.menu_opts_off(m.menu, opt.menuOptions) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) Pad() int {
	return int(C.menu_pad(m.menu))
}

func (m *Menu) Post() error {
	if C.post_menu(m.menu) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) Unpost() error {
	if C.unpost_menu(m.menu) == ERR {
		return ErrorNoInfo
	}
	return nil
}

func (m *Menu) SetCurrentItem(item *Item) error {
	if C.set_current_item(m.menu, item.item) == ERR {
		return ErrorNoInfo
	}
	return nil
}


func (i *Item) Value() bool {
	if C.item_value(i.item) == C.TRUE {
		return true
	}
	return false
}

func (i *Item) Visible() bool {
	if C.item_visible(i.item) == C.TRUE {
		return true
	}
	return false
}

func (m *Menu) Format() (int, int) {
	var crow, ccol C.int
	C.menu_format(m.menu, &crow, &ccol)
	return int(crow), int(ccol)
}
