package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
)

func Session(request *f.Request, response *f.Response, pass func()) {
	f.SessionStart(request, response, sessions.Archive)
	pass()
}
