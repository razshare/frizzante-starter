package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/indexes"
)

//go:embed .dist/*/**
var d embed.FS

func main() {
	// Create.
	s := f.ServerCreate()
	n := f.NotifierCreate()

	// Setup.
	f.ServerWithPort(s, 8080)
	f.ServerWithHostName(s, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(s, d)
	f.ServerWithNotifier(s, n)

	// Guards.
	f.ServerWithIndexGuard(s, guards.Render)
	f.ServerWithIndexGuard(s, guards.Session)

	// Routes.
	f.ServerWithIndex(s, indexes.Todos)
	f.ServerWithIndex(s, indexes.Welcome)

	// Start.
	f.ServerStart(s)
}
