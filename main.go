package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/api"
	"main/lib/guards"
	"main/lib/pages"
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
	f.ServerWithGuard(s, guards.Session)

	// Routes.
	f.ServerWithPage(s, pages.Welcome)
	f.ServerWithPage(s, pages.Todos)

	// Events.
	f.ServerWithApi(s, api.Events)

	// Start.
	f.ServerStart(s)
}
