package main

import (
	. "github.com/razshare/frizzante"
)

func main() {
	// Create server.
	server := ServerCreate()

	// Configure.
	ServerWithHostname(server, "127.0.0.1")
	ServerWithPort(server, 8080)
	ServerWithUiDirectory(server, "ui")
	ServerWithTemporaryDirectory(server, "ui/.temp")
	ServerClearTemporaryDirectory(server)

	// Route.
	ServerOnRequest(server, "GET /", func(server *Server, request *Request, response *Response) {
		SvelteComponent(response, "$pages/home", map[string]interface{}{
			"name": "world",
		})
	})

	// Log.
	ServerOnError(server, func(err error) {
		ServerLogError(server, err)
	})
	ServerOnInformation(server, func(information string) {
		ServerLogInformation(server, information)
	})

	// Start.
	ServerStart(server)
}
