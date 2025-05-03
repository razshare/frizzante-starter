package pages

import f "github.com/razshare/frizzante"

func Welcome(
	withPath f.ProvidePagePath,
	withView f.ProvidePageView,
	_ f.ProvidePageBaseHandler,
	_ f.ProvidePageActionHandler,
) {
	withPath("/")
	withView(f.ViewReference("Welcome"))
}
