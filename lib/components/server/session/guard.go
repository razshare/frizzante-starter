package session

import f "github.com/razshare/frizzante"

func pageHandler(req *f.Request, res *f.Response, _ *f.Page, pass func()) {
	f.SessionStart(req, res)
	pass()
}

func Guard(
	_ func(func(req *f.Request, res *f.Response, pass func())),
	withPageHandler func(func(req *f.Request, res *f.Response, page *f.Page, pass func())),
) {
	withPageHandler(pageHandler)
}
