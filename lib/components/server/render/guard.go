package render

import f "github.com/razshare/frizzante"

func pageHandler(_ *f.Request, _ *f.Response, page *f.Page, pass func()) {
	f.PageWithRender(page, f.RenderFull)
	pass()
}

func Guard(
	_ func(func(req *f.Request, res *f.Response, pass func())),
	withPageHandler func(func(req *f.Request, res *f.Response, page *f.Page, pass func())),
) {
	withPageHandler(pageHandler)
}
