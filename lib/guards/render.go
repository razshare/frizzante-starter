package guards

import f "github.com/razshare/frizzante"

func Render(_ *f.Request, _ *f.Response, p *f.Page, pass func()) {
	f.PageWithRender(p, f.RenderFull)
	pass()
}
