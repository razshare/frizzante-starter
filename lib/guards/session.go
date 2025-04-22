package guards

import f "github.com/razshare/frizzante"

func Session(
	withHandler func(handler func(req *f.Request, res *f.Response, pass func())),
) {
	withHandler(func(req *f.Request, res *f.Response, pass func()) {
		f.SessionStart(req, res)
		pass()
	})
}
