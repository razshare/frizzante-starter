package config

import (
	f "github.com/razshare/frizzante"
	"time"
)

func GuardSession(req *f.Request, res *f.Response) bool {
	session := f.SessionStart(req, res, SessionLoad)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
