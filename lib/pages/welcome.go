package pages

import f "github.com/razshare/frizzante"

func Welcome(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithView(page, f.ViewReference("Welcome"))
}
