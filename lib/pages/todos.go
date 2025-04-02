package pages

import (
	frz "github.com/razshare/frizzante"
	"main/lib/types"
	"strconv"
)

func Todos(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := frz.SessionStart(req, res)
	items := get("items", []types.Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	}).([]types.Item)

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
