package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guads"
	"net/url"
	"strconv"
)

func handleCheck(items []lib.Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
	items[index].Checked = true
}

func handleUncheck(items []lib.Item, form *url.Values) {
	index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
	items[index].Checked = false
}

func Todos(page *f.Page) {
	f.PageWithPath(page, "/todos")
	f.PageWithView(page, f.ViewReference("Todos"))
	f.PageWithGuardHandler(page, guads.Session)
	f.PageWithBaseHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart[lib.UserSession](request, response)

		// Inject items into view.
		f.ViewWithData(view, "items", session.Items)
	})
	f.PageWithActionHandler(page, func(request *f.Request, response *f.Response, view *f.View) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		session := f.SessionStart[lib.UserSession](request, response)

		// Read form.
		form := f.RequestReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			handleCheck(session.Items, form)
		} else if form.Has("uncheck") {
			handleUncheck(session.Items, form)
		}

		// Inject items.
		f.ViewWithData(view, "items", session.Items)
	})
}
