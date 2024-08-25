package main

import . "github.com/razshare/frizzante"

func main() {
	// Setup.
	server := ServerCreate()
	ServerWithInterface(server, "127.0.0.1")
	ServerWithPort(server, 8080)
	ServerWithTemporaryDirectory(server, ".temp")
	ServerWithNodeModulesDirectory(server, "node_modules")
	ServerWithFileServer(server, "GET /", "www")

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
