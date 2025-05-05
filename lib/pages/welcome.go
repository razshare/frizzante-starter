package pages

import f "github.com/razshare/frizzante"

func Welcome(context f.PageContext) {
	// Context.
	withpath, withView, _, _ := context()

	// Map.
	withpath("/")
	withView(f.ViewReference("Welcome"))
}
