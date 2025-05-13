package guards

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/sessions"
	"time"
)

func Session(request *f.Request, response *f.Response, pass func()) {
	session := f.SessionStart(request, response, sessions.Archive)

	if !f.SessionHas(session, "lastActivity") {
		f.SessionSetTime(session, "lastActivity", time.Now())
	}

	lastActivity := f.SessionGetTime(session, "lastActivity")

	if time.Since(lastActivity) > 30*time.Minute {
		f.SessionDestroy(session)
		f.ResponseSendNavigate(response, "Expired")
		return
	}

	if !f.SessionHas(session, "items") {
		f.SessionSetJson(session, "items", []lib.Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		})
	}

	f.SessionSetTime(session, "lastActivity", time.Now())

	pass()
}
