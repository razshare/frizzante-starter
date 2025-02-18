package pages

import (
	frz "github.com/razshare/frizzante"
	"main/schemas"
)

func Todos(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.PageConfiguration) {
	// The default session operator will destroy any session after 30 minutes of inactivity.
	get, _, _ := frz.SessionStart(req, res)
	p.Data["items"] = get("items", []schemas.Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	})
}
