package gocurse

//#cgo LDFLAGS: -lform
//#define _Bool int
//#include <form.h>
import "C"

import (
	"errors"
)

type Field struct {
	field *C.FIELD
}

type Form struct {
	form *C.FORM
}

type FieldOptions struct {
	fieldOptions C.Field_Options
}

type FormOptions struct {
	formOptions C.Form_Options
}

const (
    NO_JUSTIFICATION = C.NO_JUSTIFICATION;
    JUSTIFY_LEFT = C.JUSTIFY_LEFT;
    JUSTIFY_CENTER = C.JUSTIFY_CENTER;
    JUSTIFY_RIGHT = C.JUSTIFY_RIGHT;

    O_VISIBLE = C.O_VISIBLE;
    O_ACTIVE = C.O_ACTIVE;
    O_PUBLIC = C.O_PUBLIC;
    O_EDIT = C.O_EDIT;
    O_WRAP = C.O_WRAP;
    O_BLANK = C.O_BLANK;
    O_AUTOSKIP = C.O_AUTOSKIP;
    O_NULLOK = C.O_NULLOK;
    O_PASSOK = C.O_PASSOK;
    O_STATIC = C.O_STATIC;

    O_NL_OVERLOAD = C.O_NL_OVERLOAD;
    O_BS_OVERLOAD = C.O_BS_OVERLOAD;

    REQ_NEXT_PAGE = C.REQ_NEXT_PAGE;
    REQ_PREV_PAGE = C.REQ_PREV_PAGE;
    REQ_FIRST_PAGE = C.REQ_FIRST_PAGE;
    REQ_LAST_PAGE = C.REQ_LAST_PAGE;

    REQ_NEXT_FIELD = C.REQ_NEXT_FIELD;
    REQ_PREV_FIELD = C.REQ_PREV_FIELD;
    REQ_FIRST_FIELD = C.REQ_FIRST_FIELD;
    REQ_LAST_FIELD = C.REQ_LAST_FIELD;
    REQ_SNEXT_FIELD = C.REQ_SNEXT_FIELD;
    REQ_SPREV_FIELD = C.REQ_SPREV_FIELD;
    REQ_SFIRST_FIELD = C.REQ_SFIRST_FIELD;
    REQ_SLAST_FIELD = C.REQ_SLAST_FIELD;
    REQ_LEFT_FIELD = C.REQ_LEFT_FIELD;
    REQ_RIGHT_FIELD = C.REQ_RIGHT_FIELD;
    REQ_UP_FIELD = C.REQ_UP_FIELD;
    REQ_DOWN_FIELD = C.REQ_DOWN_FIELD;

    REQ_NEXT_CHAR = C.REQ_NEXT_CHAR;
    REQ_PREV_CHAR = C.REQ_PREV_CHAR;
    REQ_NEXT_LINE = C.REQ_NEXT_LINE;
    REQ_PREV_LINE = C.REQ_PREV_LINE;
    REQ_NEXT_WORD = C.REQ_NEXT_WORD;
    REQ_PREV_WORD = C.REQ_PREV_WORD;
    REQ_BEG_FIELD = C.REQ_BEG_FIELD;
    REQ_END_FIELD = C.REQ_END_FIELD;
    REQ_BEG_LINE = C.REQ_BEG_LINE;
    REQ_END_LINE = C.REQ_END_LINE;
    REQ_LEFT_CHAR = C.REQ_LEFT_CHAR;
    REQ_RIGHT_CHAR = C.REQ_RIGHT_CHAR;
    REQ_UP_CHAR = C.REQ_UP_CHAR;
    REQ_DOWN_CHAR = C.REQ_DOWN_CHAR;

    REQ_NEW_LINE = C.REQ_NEW_LINE;
    REQ_INS_CHAR = C.REQ_INS_CHAR;
    REQ_INS_LINE = C.REQ_INS_LINE;
    REQ_DEL_CHAR = C.REQ_DEL_CHAR;
    REQ_DEL_PREV = C.REQ_DEL_PREV;
    REQ_DEL_LINE = C.REQ_DEL_LINE;
    REQ_DEL_WORD = C.REQ_DEL_WORD;
    REQ_CLR_EOL = C.REQ_CLR_EOL;
    REQ_CLR_EOF = C.REQ_CLR_EOF;
    REQ_CLR_FIELD = C.REQ_CLR_FIELD;
    REQ_OVL_MODE = C.REQ_OVL_MODE;
    REQ_INS_MODE = C.REQ_INS_MODE;
    REQ_SCR_FLINE = C.REQ_SCR_FLINE;
    REQ_SCR_BLINE = C.REQ_SCR_BLINE;
    REQ_SCR_FPAGE = C.REQ_SCR_FPAGE;
    REQ_SCR_BPAGE = C.REQ_SCR_BPAGE;
    REQ_SCR_FHPAGE = C.REQ_SCR_FHPAGE;
    REQ_SCR_BHPAGE = C.REQ_SCR_BHPAGE;
    REQ_SCR_FCHAR = C.REQ_SCR_FCHAR;
    REQ_SCR_BCHAR = C.REQ_SCR_BCHAR;
    REQ_SCR_HFLINE = C.REQ_SCR_HFLINE;
    REQ_SCR_HBLINE = C.REQ_SCR_HBLINE;
    REQ_SCR_HFHALF = C.REQ_SCR_HFHALF;
    REQ_SCR_HBHALF = C.REQ_SCR_HBHALF;

    REQ_VALIDATION = C.REQ_VALIDATION;
    REQ_NEXT_CHOICE = C.REQ_NEXT_CHOICE;
    REQ_PREV_CHOICE = C.REQ_PREV_CHOICE;

    MIN_FORM_COMMAND = C.MIN_FORM_COMMAND;
    MAX_FORM_COMMAND = C.MAX_FORM_COMMAND;
)

var (
	ErrorSystemError = errors.New("System error")
	ErrorBadArgument = errors.New("bad argument error")
	ErrorPosted = errors.New("Posted error")
	ErrorConnected = errors.New("connected error")
	ErrorBadState = errors.New("bad state error")
	ErrorNoRoom = errors.New("no room error")
	ErrorNotPosted = errors.New("not posted error")
	ErrorUnknownCommand = errors.New("unkown command error")
	ErrorNoMatch = errors.New("no match error")
	ErrorNotSelectable = errors.New("not selectable error")
	ErrorNotConnected = errors.New("not connected error")
	ErrorRequestDenied = errors.New("request denied error")
	ErrorInvalidField = errors.New("invalid field error")
	ErrorCurrent = errors.New("current error")
)

func geterror(errno C.int) error {
	switch errno {
	case 0: return nil
	case -1: return ErrorSystemError
	case -2: return ErrorBadArgument
	case -3: return ErrorPosted
	case -4: return ErrorConnected
	case -5: return ErrorBadState
	case -6: return ErrorNoRoom
	case -7: return ErrorNotPosted
	case -8: return ErrorUnknownCommand
	case -9: return ErrorNoMatch
	case -10: return ErrorNotSelectable
	case -11: return ErrorNotConnected
	case -12: return ErrorRequestDenied
	case -13: return ErrorInvalidField
	case -14: return ErrorCurrent
	}
	return Error
}


/*
* FIELD METHODS
 */

func NewField(height int, width int, top int, left int, offscreen int, nbuf int) (*Field, error) {
	field := C.new_field(C.int(height), C.int(width), C.int(top), C.int(left), C.int(offscreen), C.int(nbuf))
	if field == nil {
		return nil, errors.New("NewField failed")
	}
	return &Field{field}, nil
}

func (f *Field) DupField(top int, left int) (*Field, error) {
	dup := C.dup_field(f.field, C.int(top), C.int(left))
	if dup == nil {
		return nil, errors.New("Field.Dup failed")
	}
	return &Field{dup}, nil
}

func (f *Field) Link(top int, left int) (*Field, error) {
	link := C.link_field(f.field, C.int(top), C.int(left))
	if link == nil {
		return nil, errors.New("Field.Link failed")
	}
	return &Field{link}, nil
}

func (f *Field) Free() error {
	if C.free_field(f.field) != OK {
		return errors.New("Field.Free failed")
	}
	return nil
}

func (f *Field) Info() (int, int, int, int, int, int, error) {
	var (
		height    C.int
		width     C.int
		top       C.int
		left      C.int
		offscreen C.int
		nbuf      C.int
	)
	
	if C.field_info(f.field, &height, &width, &top, &left, &offscreen, &nbuf) == ERR {
		return 0, 0, 0, 0, 0, 0, errors.New("Field.Info failed")
	}
	return int(height), int(width), int(top), int(left), int(offscreen), int(nbuf), nil
}

func (f *Field) DynamicInfo() (int, int, int, error) {
	var (
		prows C.int
		pcols C.int
		pmax  C.int
	)
	if C.dynamic_field_info(f.field, &prows, &pcols, &pmax) == ERR {
		return 0, 0, 0, errors.New("Field.DynamicInfo failed")
	}
	return int(prows), int(pcols), int(pmax), nil
}

func (f *Field) SetMax(max int) error {
	return geterror(C.set_max_field(f.field, C.int(max)))

}

func (f *Field) Move(x int, y int) error {
	return geterror(C.move_field(f.field, C.int(x), C.int(y)))
}

func (f *Field) SetNewPage(newPage bool) error {
	return geterror(C.set_new_page(f.field, boolToCint(newPage)))
}

func (f *Field) SetJust(justMode int) error {
	return geterror(C.set_field_just(f.field, C.int(justMode)))
}

func (f *Field) Just() int {
	return int(C.field_just(f.field))
}

func (f *Field) SetFore(fore int) error {
	return geterror(C.set_field_fore(f.field, C.chtype(fore)))
}

func (f *Field) SetBack(back int) error {
	return geterror(C.set_field_back(f.field, C.chtype(back)))
}

func (f *Field) SetPad(pad int) error {
	return geterror(C.set_field_pad(f.field, C.int(pad)))
}

func (f *Field) Pad() int {
	return int(C.field_pad(f.field))
}

func (f *Field) SetBuffer(ind int, message string) error { //Memoryleak
	return geterror(C.set_field_buffer(f.field, C.int(ind), C.CString(message)))
}

func (f *Field) SetStatus(status bool) error {
	return geterror(C.set_field_status(f.field, boolToCint(status)))
}

func (f *Field) SetOpts(attr FieldOptions) error {
	return geterror(C.set_field_opts(f.field, attr.fieldOptions))
}

func (f *Field) OptsOn(attr FieldOptions) error {
	return geterror(C.field_opts_on(f.field, attr.fieldOptions))
}

func (f *Field) OptsOff(attr FieldOptions) error {
	return geterror(C.field_opts_off(f.field, attr.fieldOptions))
}

func (f *Field) Buffer(ind int) string {
	buf := C.field_buffer(f.field, C.int(ind))
	return C.GoString(buf)
}

func (f *Field) Fore() int {
	return int(C.field_fore(f.field))
}

func (f *Field) Back() int {
	return int(C.field_back(f.field))
}

func (f *Field) NewPage() bool {
	return intToBool(C.new_page(f.field))
}

func (f *Field) Opts() FieldOptions {
	return FieldOptions{C.field_opts(f.field)}
}

func (f *Field) Index() int {
    return int(C.field_index(f.field))
}

/*
* FORM METHODS
 */

func NewForm(fields []*Field) (*Form, error) {
	f := make([]*C.FIELD, 0, len(fields))
	for _, value := range fields {
		f = append(f, value.field)
	}
	form := C.new_form(&f[0])
	if form == nil {
		return nil, errors.New("NewForm failed")
	}
	return &Form{form}, nil
}

func (f *Form) CurrentField() *Field {
	return &Field{C.current_field(f.form)}
}

func (f *Form) DataAhead() bool {
    return intToBool(C.data_ahead(f.form))
}

func (f *Form) DataBehind() bool {
    return intToBool(C.data_behind(f.form))
}

func (f *Form) Free() error {
	return geterror(C.free_form(f.form))
}

func (f *Form) SetFields(fields []*Field) error {
	fs := make([]*C.FIELD, 0, len(fields))
	for _, value := range fields {
		fs = append(fs, value.field)
	}
	return geterror(C.set_form_fields(f.form, &fs[0]))
}

func (f *Form) FieldCount() int {
	return int(C.field_count(f.form))
}

func (f *Form) SetCurrentField(field *Field) error {
	return geterror(C.set_current_field(f.form, field.field))
}

func (f *Form) SetPage(page int) error {
	return geterror(C.set_form_page(f.form, C.int(page)))
}

func (f *Form) Page() int {
	return int(C.form_page(f.form))
}

func (f *Form) Post() error {
	return geterror(C.post_form(f.form))
}

func (f *Form) Unpost() error {
	return geterror(C.unpost_form(f.form))
}

func (f *Form) Drive(req int) error {
	return geterror(C.form_driver(f.form, C.int(req)))
}

func (f *Form) Opts() FormOptions {
    return FormOptions{C.form_opts(f.form)}
}

func (f *Form) SetOpts(attr FormOptions) error {
	return geterror(C.set_form_opts(f.form, attr.formOptions))
}

func (f *Form) OptsOn(attr FormOptions) error {
	return geterror(C.form_opts_on(f.form, attr.formOptions))
}

func (f *Form) OptsOff(attr FormOptions) error {
	return geterror(C.form_opts_off(f.form, attr.formOptions))
}

func (f *Form) Scale() (int, int, error) {
	var (
		rows C.int
		cols C.int
	)
	if C.scale_form(f.form, &rows, &cols) != C.OK {
		return 0, 0, errors.New("Form.Scale failed")
	}
	return int(rows), int(cols), nil
}

func (f *Form) SetWin(window *Window) error {
	return geterror(C.set_form_win(f.form, window.win))
}

func (f *Form) SetSub(window *Window) error {
	return geterror(C.set_form_sub(f.form, window.win))
}

func (f *Form) Win() *Window {
	return &Window{C.form_win(f.form)}
}

func (f *Form) Sub() *Window {
	return &Window{C.form_sub(f.form)}
} 
