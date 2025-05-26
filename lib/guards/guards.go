package guards

import (
	"github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func NotExpired(req *frizzante.Request, res *frizzante.Response) bool {
	session := frizzante.SessionStart(req, res, sessions.Adapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("expired")
		return false
	}

	session.Data.LastActivity = time.Now()
	return true
}
