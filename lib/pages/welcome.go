package pages

import f "github.com/razshare/frizzante"

func Welcome(context f.PageContext) {
	// Context.
	path, view, _, _ := context()

	// Map.
	path("/")
	view(f.ViewReference("Welcome"))
}
