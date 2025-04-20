package todos

import (
	f "github.com/razshare/frizzante"
	"net/url"
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

func baseHandler(req *f.Request, res *f.Response, p *f.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := f.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	f.PageWithData(p, "items", items)
}

func actionHandler(req *f.Request, res *f.Response, p *f.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := f.SessionStart(req, res)
	items := get("items", initialItems).([]Item)

	// Read form.
	form := f.ReceiveForm(req)

	// Handle checks.
	if form.Has("check") {
		handleCheck(items, form)
	} else if form.Has("uncheck") {
		handleUncheck(items, form)
	}

	f.PageWithData(p, "items", items)
}

func handleCheck(items []Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
	items[index].Checked = true
}

func handleUncheck(items []Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
	items[index].Checked = false
}

func Index(
	withPage func(page string),
	withPath func(path string),
	withBaseHandler func(baseHandler func(req *f.Request, res *f.Response, page *f.Page)),
	withActionHandler func(actionFunction func(req *f.Request, res *f.Response, page *f.Page)),
) {
	withPage("todos")
	withPath("/todos")
	withBaseHandler(baseHandler)
	withActionHandler(actionHandler)
}
