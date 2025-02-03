package pages

import (
	frz "github.com/razshare/frizzante"
	"main/schemas"
)

func Todos(_ *frz.Server, req *frz.Request, res *frz.Response) *frz.SveltePageConfiguration {
	get, _ := frz.SessionStart(req, res)

	items := get("items", []schemas.Item{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	})

	return &frz.SveltePageConfiguration{
		Render: frz.ModeFull,
		Data: map[string]interface{}{
			"items": items,
		},
	}
}
