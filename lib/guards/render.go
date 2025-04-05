package guards

import frz "github.com/razshare/frizzante"

func Render(_ *frz.Request, _ *frz.Response, p *frz.Page, pass func()) {
	frz.PageWithRender(p, frz.RenderFull)
	pass()
}
