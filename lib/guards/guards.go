package guards

import (
	"github.com/razshare/frizzante"
	"main/lib/sessions"
	"time"
)

func SessionIsValid(req *frizzante.Request, res *frizzante.Response, pass func()) {
	session := frizzante.SessionStart(req, res, sessions.Adapter)

	if time.Since(session.Data.LastActivity) > 30*time.Minute {
		session.Destroy()
		res.SendNavigate("expired")
		return
	}

	session.Data.LastActivity = time.Now()
	pass()
}
