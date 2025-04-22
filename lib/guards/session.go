package guards

import f "github.com/razshare/frizzante"

func Session(
	withHandler func(handler func(request *f.Request, response *f.Response, pass func())),
) {
	withHandler(func(request *f.Request, response *f.Response, pass func()) {
		f.SessionStart(request, response)
		pass()
	})
}
