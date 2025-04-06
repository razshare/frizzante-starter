package guards

import f "github.com/razshare/frizzante"

func Session(req *f.Request, res *f.Response, _ *f.Page, pass func()) {
	f.SessionStart(req, res)
	pass()
}
