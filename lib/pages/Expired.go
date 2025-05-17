package pages

import f "github.com/razshare/frizzante"

type ExpiredData struct{}

func Expired(page *f.Page[ExpiredData]) {
	// Configure.
	f.PageWithPath(page, "/Expired")
	f.PageWithName(page, "Expired")
	f.PageWithView(page, "Expired", func() ExpiredData {
		return ExpiredData{}
	})

	// Guard.
	// Noop.

	// Handle base.
	// Noop.

	// Handle action.
	// Noop.
}
