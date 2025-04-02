package api

import (
	frz "github.com/razshare/frizzante"
)

func Check(_ *frz.Server, req *frz.Request, res *frz.Response) {
	if !frz.VerifyContentType(req, "application/json") {
		frz.SendStatus(res, 400)
		return
	}

	get, set, _ := frz.SessionStart(req, res)
	items := get("items", []any{})
	frz.ReceiveJson(req, &items)
	set("items", items)
}
