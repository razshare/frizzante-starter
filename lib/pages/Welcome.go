package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
)

type WelcomeData struct {
	Items []lib.Item
}

func Welcome(page *f.Page[WelcomeData]) {
	// Configure.
	f.PageWithPath(page, "/")
	f.PageWithName(page, "Welcome")
	f.PageWithView(page, "Welcome", func() WelcomeData {
		return WelcomeData{}
	})

	// Guard.
	f.PageWithGuardHandler(page, guards.Session)

	// Handle base.
	// Noop.

	// Handle action.
	// Noop.
}
