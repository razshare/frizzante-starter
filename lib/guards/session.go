package guards

import f "github.com/razshare/frizzante"

func Session(context f.GuardContext) {
	// Context.
	handler := context()

	// Configure.
	handler(func(request *f.Request, response *f.Response, pass func()) {
		// Start session.
		f.SessionStart(request, response)
		pass()
	})
}
