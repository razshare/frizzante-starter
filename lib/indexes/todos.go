package indexes

import (
	frz "github.com/razshare/frizzante"
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

func showTodos(req *frz.Request, res *frz.Response, p *frz.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := frz.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	frz.PageWithData(p, "items", items)
}

func updateTodos(req *frz.Request, res *frz.Response, p *frz.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := frz.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	// Read form.
	form := frz.ReceiveForm(req)
	if form.Has("check") {
		index, pe := strconv.ParseInt(form.Get("check"), 10, 32)
		if nil != pe {
			frz.PageWithData(p, "error", pe.Error())
			return
		}
		items[index].Checked = true
	} else if form.Has("uncheck") {
		index, pe := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		if nil != pe {
			frz.PageWithData(p, "error", pe.Error())
			return
		}
		items[index].Checked = false
	}

	frz.PageWithData(p, "items", items)
}

func Todos() (show frz.PageFunction, action frz.PageFunction) {
	show = showTodos
	action = updateTodos
	return
}
