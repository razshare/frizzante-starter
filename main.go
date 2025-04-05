package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/indexes"
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

	// Guards.
	frz.ServerWithPageGuard(s, guards.Render)
	frz.ServerWithPageGuard(s, guards.Session)

	// Routes.
	frz.ServerWithPage(s, "/todos", "todos", indexes.Todos)
	frz.ServerWithPage(s, "/", "welcome", indexes.Welcome)

	// Start.
	frz.ServerStart(s)
}
