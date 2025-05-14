package pages

import f "github.com/razshare/frizzante"

func Expired(page *f.Page) {
	f.PageWithPath(page, "/Expired")
	f.PageWithView(page, f.ViewReference("Expired"))
}
