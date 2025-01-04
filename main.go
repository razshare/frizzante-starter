package main

import frz "github.com/razshare/frizzante"

func main() {
	// Create server.
	server := frz.ServerCreate()

	// Configure.
	frz.ServerWithHostname(server, "127.0.0.1")
	frz.ServerWithPort(server, 8080)
	frz.ServerWithUiDirectory(server, "ui")
	frz.ServerWithTemporaryDirectory(server, "ui/.temp")
	frz.ServerClearTemporaryDirectory(server)

	// Route.
	frz.ServerOnRequest(server, "GET /", func(server *frz.Server, request *frz.Request, response *frz.Response) {
		frz.SvelteComponent(response, "$pages/home", map[string]interface{}{
			"name": "world",
		})
	})
	frz.ServerOnRequest(server, "GET /hello", func(Server *frz.Server, request *frz.Request, response *frz.Response) {
		frz.Echo(response, "hello")
	})

	// Log.
	frz.ServerOnError(server, func(err error) {
		frz.ServerLogError(server, err)
	})
	frz.ServerOnInformation(server, func(information string) {
		frz.ServerLogInformation(server, information)
	})

	// Start.
	frz.ServerStart(server)
}
