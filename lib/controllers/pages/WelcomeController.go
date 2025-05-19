package pages

import (
	f "github.com/razshare/frizzante"
)

type WelcomeController struct {
	f.PageController
}

func (_ WelcomeController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path:         "/",
		TryFileFirst: true,
	}
}

func (_ WelcomeController) Base(_ *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}

func (_ WelcomeController) Action(_ *f.Request, res *f.Response) {
	res.SendView(f.NewView(f.RenderModeFull))
}
