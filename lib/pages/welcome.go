package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath func(string),
	withView func(*f.View),
	_ func(func(*f.Request, *f.Response, *f.View)),
	_ func(func(*f.Request, *f.Response, *f.View)),
) {
	withPath("/")
	withView(f.ViewReference("Welcome"))
}
