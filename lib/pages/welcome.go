package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath f.WithPagePath,
	withView f.WithPageView,
	_ f.WithPageBaseHandler,
	_ f.WithPageActionHandler,
) {
	withPath("/")
	withView(f.ViewReference("Welcome"))
}
