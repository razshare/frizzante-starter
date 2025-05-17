package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func Session(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response, sessions.Archived)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		f.SessionDestroy(session)
		f.ResponseSendNavigate(response, "Expired")
		return
	}

	session.Data.LastActivity = time.Now()

	pass()
}
