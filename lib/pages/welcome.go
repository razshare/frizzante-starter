package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath func(path string),
	withView func(view *f.View),
	withBaseHandler func(baseHandler func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	)),
	_ func(actionFunction func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	)),
) {
	withPath("/")
	withView(f.ViewReference("Welcome"))
	withBaseHandler(func(
		request *f.Request,
		response *f.Response,
		view *f.View,
	) {
		view.Render = f.RenderClient
	})
}
