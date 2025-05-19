package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func SessionIsValid(req *f.Request, res *f.Response) bool {
	session := f.SessionStart(req, res, sessions.Archived)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("Expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
