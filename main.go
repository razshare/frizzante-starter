package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"log"
	"main/pages"
	"main/routes"
)

//go:embed .dist/*/**
var efs embed.FS

// Noop.
func noop(
	_ *frz.Server,
	_ *frz.Request,
	_ *frz.Response,
	_ *frz.PageConfiguration,
) {
}

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
	frz.ServerWithRoute(srv, "POST /check", frz.Route(routes.Check))
	frz.ServerWithRoute(srv, "GET /todos", frz.Page("todos", pages.Todos))
	frz.ServerWithRoute(srv, "GET /", frz.Page("welcome", noop))

	// Log.
	frz.ServerWithErrorReceiver(srv, func(err error) {
		logger.Fatal(err)
	})

	// Start.
	frz.ServerStart(srv)
}
