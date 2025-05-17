package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/sessions"
	"net/url"
	"strconv"
)

func check(items []lib.Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
	items[index].Checked = true
}

func uncheck(items []lib.Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
	items[index].Checked = false
}

type TodosData struct {
	Items []lib.Item `json:"items"`
}

func Todos(page *f.Page[TodosData]) {
	// Configure.
	f.PageWithPath(page, "/Todos")
	f.PageWithName(page, "Todos")
	f.PageWithView(page, "Todos", func() TodosData {
		return TodosData{}
	})

	// Guard.
	f.PageWithGuardHandler(page, guards.Session)

	// Handle base.
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, view *f.View[TodosData]) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response, sessions.Archived)

		// Inject items into view.
		view.Data.Items = session.Data.Items
	})

	// Handle action.
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View[TodosData]) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response, sessions.Archived)

		// Get items.
		items := session.Data.Items

		// Read form.
		form := f.RequestReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			check(items, form)
		} else if form.Has("uncheck") {
			uncheck(items, form)
		}

		// Update view items.
		view.Data.Items = items

		// Update session items.
		session.Data.Items = items

		// Save session.
		f.SessionSave(session)
	})
}
