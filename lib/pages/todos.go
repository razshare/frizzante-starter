package pages

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

func handleCheck(items []Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
	items[index].Checked = true
}

func handleUncheck(items []Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
	items[index].Checked = false
}

func Todos(
	withPath func(path string),
	withDocument func(doc *f.Document),
	withBaseHandler func(baseHandler func(req *f.Request, res *f.Response, doc *f.Document)),
	withActionHandler func(actionFunction func(req *f.Request, res *f.Response, doc *f.Document)),
) {
	withPath("/todos")
	withDocument(f.DocumentCreate("todos"))

	// Base.
	withBaseHandler(func(req *f.Request, res *f.Response, doc *f.Document) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		get, _, _ := f.SessionStart(req, res)
		doc.Data["items"] = get("items", initialItems)
	})

	// Action.
	withActionHandler(func(req *f.Request, res *f.Response, doc *f.Document) {
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

		doc.Data["items"] = items
	})
}
