package guards

import frz "github.com/razshare/frizzante"

func Session(req *frz.Request, res *frz.Response, _ *frz.Page, pass func()) {
	frz.SessionStart(req, res)
	pass()
}
