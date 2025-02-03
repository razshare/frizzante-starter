package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/handlers"
	"main/pages"
)

//go:embed .dist/*/**
var efs embed.FS

func main() {
	// Create.
	srv := frz.ServerCreate()

	// Setup.
	frz.ServerWithPort(srv, 8080)
	frz.ServerWithHostName(srv, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(srv, efs)

	// Route (order matters, "/" should always be last).
	frz.ServerWithSveltePage(srv, "GET /todos", "todos", pages.Todos)
	frz.ServerWithRequestHandler(srv, "POST /check", handlers.Check)
	frz.ServerWithSveltePage(srv, "GET /", "welcome", pages.Welcome)

	// Log.
	frz.ServerOnError(srv, func(err error) {
		frz.ServerLogError(srv, err)
	})
	frz.ServerOnInformation(srv, func(inf string) {
		frz.ServerLogInformation(srv, inf)
	})

	// Start.
	frz.ServerStart(srv)
}
