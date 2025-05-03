package guards

import f "github.com/razshare/frizzante"

func Session(withHandler func(func(*f.Request, *f.Response, func()))) {
	// Guard.
	withHandler(func(request *f.Request, response *f.Response, pass func()) {
		// Start session.
		f.SessionStart(request, response)
		pass()
	})
}
