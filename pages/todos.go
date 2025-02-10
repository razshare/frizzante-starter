package pages

import (
	frz "github.com/razshare/frizzante"
	"main/schemas"
)

func Todos(_ *frz.Server, req *frz.Request, res *frz.Response, p *frz.Page) {
	get, _, _ := frz.SessionStart(req, res)

	items := get("items", []schemas.Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	})

	frz.PageWithRenderMode(p, frz.ModeFull)
	frz.PageWithData(p, "items", items)
}
