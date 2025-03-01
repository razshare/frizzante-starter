package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"log"
	"main/api"
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
	frz.ServerWithLogger(srv, logger)

	// Route (order matters, "/" should always be last).
	frz.ServerRouteApi(srv, "POST /check", api.Check)
	frz.ServerRoutePage(srv, "GET /todos", "todos", pages.Todos)
	frz.ServerRoutePage(srv, "GET /", "welcome", pages.Welcome)

	// Log.
	frz.ServerRecallError(srv, func(err error) {
		logger.Fatal(err)
	})

	// Start.
	frz.ServerStart(srv)
}
