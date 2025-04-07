package indexes

import (
	f "github.com/razshare/frizzante"
	"strconv"
)

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

var initialItems = []Item{
	{Checked: false, Description: "Pet the cat."},
	{Checked: false, Description: "Do laundry"},
	{Checked: false, Description: "Pet the cat."},
	{Checked: false, Description: "Cook"},
	{Checked: false, Description: "Pet the cat."},
}

func todosShow(req *f.Request, res *f.Response, p *f.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := f.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	f.PageWithData(p, "items", items)
}

func todosAction(req *f.Request, res *f.Response, p *f.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := f.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	// Read form.
	form := f.ReceiveForm(req)
	if form.Has("check") {
		index, pe := strconv.ParseInt(form.Get("check"), 10, 32)
		if nil != pe {
			f.PageWithData(p, "error", pe.Error())
			return
		}
		items[index].Checked = true
	} else if form.Has("uncheck") {
		index, pe := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		if nil != pe {
			f.PageWithData(p, "error", pe.Error())
			return
		}
		items[index].Checked = false
	}

	f.PageWithData(p, "items", items)
}

func Todos() (
	page string,
	show f.PageFunction,
	action f.PageFunction,
) {
	page = "todos"
	show = todosShow
	action = todosAction
	return
}
