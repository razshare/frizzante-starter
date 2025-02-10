package pages

import frz "github.com/razshare/frizzante"

func Welcome(_ *frz.Server, _ *frz.Request, _ *frz.Response, p *frz.Page) {
	frz.PageWithRenderMode(p, frz.ModeFull)
}
