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

func Todos(page *f.Page) {
	f.PageWithPath(page, "/todos")
	f.PageWithView(page, f.ViewReference("Todos"))
	f.PageWithBase(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)
		items := f.SessionGet[[]Item](session, "items", initialItems)
		f.ViewWithData(view, "items", items)
	})
	f.PageWithAction(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)
		items := f.SessionGet[[]Item](session, "items", initialItems)

		// Read form.
		form := f.RequestReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			handleCheck(items, form)
		} else if form.Has("uncheck") {
			handleUncheck(items, form)
		}
		f.ViewWithData(view, "items", items)
	})
}
