package guards

import f "github.com/razshare/frizzante"

func Session(context f.GuardContext) {
	// Context.
	withHandler := context()

	// Configure.
	withHandler(func(request *f.Request, response *f.Response, pass func()) {
		// Start session.
		f.SessionStart(request, response)
		pass()
	})
}
