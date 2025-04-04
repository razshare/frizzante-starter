package pages

import (
	frz "github.com/razshare/frizzante"
	"main/lib"
	"strconv"
)

func Todos(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := frz.SessionStart(req, res)
	items := get("Items", []lib.Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	}).([]lib.Item)

	// Read form.
	form := frz.ReceiveForm(req)
	if form.Has("Check") {
		index, pe := strconv.ParseInt(form.Get("Check"), 10, 32)
		if nil != pe {
			frz.PageWithData(p, "Error", pe.Error())
			return
		}
		items[index].Checked = true
	} else if form.Has("Uncheck") {
		index, pe := strconv.ParseInt(form.Get("Uncheck"), 10, 32)
		if nil != pe {
			frz.PageWithData(p, "Error", pe.Error())
			return
		}
		items[index].Checked = false
	}

	frz.PageWithData(p, "Items", items)
}
