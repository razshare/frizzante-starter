package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"log"
	"main/handlers"
	"main/pages"
)

//go:embed .dist/*/**
var efs embed.FS

func main() {
	// Create.
	logger := log.Default()
	srv := frz.ServerCreate()

	// Setup.
	frz.ServerWithPort(srv, 8080)
	frz.ServerWithHostName(srv, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(srv, efs)
	frz.ServerWithSessionHandler(srv, handlers.Session)
	frz.ServerWithLogger(srv, logger)

	// Route (order matters, "/" should always be last).
	frz.ServerWithRequestHandler(srv, "POST /check", handlers.Check)
	frz.ServerWithSveltePage(srv, "GET /todos", "todos", pages.Todos)
	frz.ServerWithSveltePage(srv, "GET /", "welcome", pages.Welcome)

	// Log.
	frz.ServerWithErrorHandler(srv, func(err error) {
		logger.Fatal(err)
	})

	// Start.
	frz.ServerStart(srv)
}
