package pages

import f "github.com/razshare/frizzante"

func Expired(page *f.Page) {
	f.PageWithPath(page, "/expired")
	f.PageWithView(page, f.ViewReference("Expired"))
}
