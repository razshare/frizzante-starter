package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guards"
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
	f.PageWithGuardHandler(page, guards.Session)
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)

		// Make sure "items" exists.
		if !f.SessionHas(session, "items") {
			f.SessionSet(session, "items", initialItems)
		}

		// Retrieve and inject items.
		items := f.SessionGet[[]Item](session, "items")
		f.ViewWithData(view, "items", items)
	})
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)

		// Make sure "items" exists.
		if !f.SessionHas(session, "items") {
			f.SessionSet(session, "items", initialItems)
		}

		// Retrieve items.
		items := f.SessionGet[[]Item](session, "items")

		// Read form.
		form := f.RequestReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			handleCheck(items, form)
		} else if form.Has("uncheck") {
			handleUncheck(items, form)
		}

		// PUpdate session.
		f.SessionSet(session, "items", items)

		// Inject items.
		f.ViewWithData(view, "items", items)
	})
}
