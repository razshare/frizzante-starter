package guads

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

func Session(request *f.Request, response *f.Response, pass func()) {
	f.SessionStart[lib.UserSession](request, response)
	pass()
}
