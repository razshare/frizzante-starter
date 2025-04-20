package session

import f "github.com/razshare/frizzante"

func guardHandler(req *f.Request, res *f.Response, pass func()) {
	f.SessionStart(req, res)
	pass()
}

func Guard(
	withGuardHandler func(
		guardHandler func(
			req *f.Request,
			res *f.Response,
			pass func(),
		),
	),
) {
	withGuardHandler(guardHandler)
}
