package main

import (
	. "github.com/razshare/frizzante"
)

func main() {
	// Setup.
	server := ServerCreate()
	ServerWithHostname(server, "127.0.0.1")
	ServerWithPort(server, 8080)
	ServerWithTemporaryDirectory(server, ".temp")
	ServerClearTemporaryDirectory(server)

	// Routes.
	ServerOnRequest(server, "GET /", func(server *Server, request *Request, response *Response) {
		Svelte(response, "/index", "<h3>hello</h3>")
	})

	// Logging.
	ServerOnError(server, func(err error) {
		ServerLogError(server, err)
	})
	ServerOnInformation(server, func(information string) {
		ServerLogInformation(server, information)
	})

	// Start.
	startError := ServerStart(server)
	if startError != nil {
		println(startError.Error())
		return
	}
}
