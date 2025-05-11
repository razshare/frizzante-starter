package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guads"
)

func Welcome(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithView(page, f.ViewReference("Welcome"))
	f.PageWithGuardHandler(page, guads.Session)
}
