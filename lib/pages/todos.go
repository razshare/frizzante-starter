package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
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

func Todos(page *f.Page) {
	f.PageWithPath(page, "/todos")
	f.PageWithView(page, f.ViewReference("Todos"))
	f.PageWithGuardHandler(page, guards.Session)
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)

		// Get items.
		items := f.SessionGetJson[[]lib.Item](session, "items")

		// Inject items into view.
		f.ViewWithData(view, "items", items)
	})
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart(request, response)

		// Get items.
		items := f.SessionGetJson[[]lib.Item](session, "items")

		// Read form.
		form := f.RequestReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			check(items, form)
		} else if form.Has("uncheck") {
			uncheck(items, form)
		}

		// Inject items.
		f.ViewWithData(view, "items", items)

		// Update session.
		f.SessionSetJson(session, "items", items)
	})
}
