package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func SessionIsValid(request *f.Request, response *f.Response) bool {
	session := f.SessionStart(request, response, sessions.Archived)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		response.SendNavigate("Expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
