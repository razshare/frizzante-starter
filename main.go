package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib/api"
	"main/lib/pages"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	// Create.
	s := frz.ServerCreate()
	n := frz.NotifierCreate()

	// Setup.
	frz.ServerWithPort(s, 8080)
	frz.ServerWithHostName(s, "127.0.0.1")
	frz.ServerWithEmbeddedFileSystem(s, dist)
	frz.ServerWithNotifier(s, n)

	// Route (order matters, "/" should always be last).
	frz.ServerRouteApi(s, "POST /check", api.Check)
	frz.ServerRoutePage(s, "GET /todos", "todos", pages.Todos)
	frz.ServerRoutePage(s, "POST /todos", "todos", pages.Todos)
	frz.ServerRoutePage(s, "GET /", "welcome", pages.Welcome)

	// Start.
	frz.ServerStart(s)
}
