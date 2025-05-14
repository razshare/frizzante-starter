package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib/guards"
)

func Welcome(page *f.Page) {
	f.PageWithPath(page, "/")
	f.PageWithView(page, f.ViewReference("Welcome"))
	f.PageWithGuardHandler(page, guards.Session)
}
