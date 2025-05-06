package pages

import f "github.com/razshare/frizzante"

func Welcome(context f.PageContext) {
	// Context.
	withPath, withView, _, _ := context()

	// Map.
	withPath("/")
	withView(f.ViewReference("Welcome"))
}
