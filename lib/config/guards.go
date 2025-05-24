package config

import (
	"time"

	f "github.com/razshare/frizzante"
)

func GuardNotExpired(req *f.Request, res *f.Response) bool {
	session := f.SessionStart(req, res, SessionAdapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
