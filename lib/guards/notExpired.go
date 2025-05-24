package guards

import (
	"main/lib/sessions"
	"time"

	f "github.com/razshare/frizzante"
)

func NotExpired(req *f.Request, res *f.Response) bool {
	session := f.SessionStart(req, res, sessions.Adapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
