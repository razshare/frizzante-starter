package guards

import f "github.com/razshare/frizzante"

func Session(request *f.Request, response *f.Response, pass func()) {
	// Start session.
	f.SessionStart(request, response)
	pass()
}
