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
	withView func(view *f.View),
	withBaseHandler func(baseHandler func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	)),
	withActionHandler func(actionFunction func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	)),
) {
	withPath("/todos")
	withView(f.ViewReference("Todos"))

	// Base.
	withBaseHandler(func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		get, _, _ := f.SessionStart(request, response)
		f.ViewWithData(view, "items", get("items", initialItems))
	})

	// Action.
	withActionHandler(func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	) {
		// The default session operator will destroy any session after 30 minutes of inactivity.
		get, _, _ := f.SessionStart(request, response)
		items := get("items", initialItems).([]Item)

		// Read form.
		form := f.ReceiveForm(request)

		// Handle checks.
		if form.Has("check") {
			handleCheck(items, form)
		} else if form.Has("uncheck") {
			handleUncheck(items, form)
		}
		f.ViewWithData(view, "items", items)
	})
}
