package pages

import f "github.com/razshare/frizzante"

type ExpiredController struct {
	f.PageController
}

func (_ ExpiredController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/expired",
	}
}

func (_ ExpiredController) Base(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}

func (_ ExpiredController) Action(_ *f.Request, response *f.Response) {
	response.SendView(f.NewView(f.RenderModeFull))
}
