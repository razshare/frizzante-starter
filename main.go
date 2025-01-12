package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
)

//go:embed www/dist/*/**
var efs embed.FS

func main() {
	// Create.
	server := frz.ServerCreate()

	// Configure.
	frz.ServerWithPort(server, 8080)
	frz.ServerWithHostName(server, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(server, efs)
	frz.ServerClearTemporaryDirectory(server)

	// Route.
	frz.ServerSetSveltePage(server, true, "GET /about", "about", nil)
	frz.ServerSetSveltePage(server, true, "GET /", "welcome", nil)

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
