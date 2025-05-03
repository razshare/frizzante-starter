package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/api"
	"main/lib/guards"
	"main/lib/pages"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	// Create.
	server := f.ServerCreate()
	notifier := f.NotifierCreate()

	// Setup.
	f.ServerWithPort(server, 8080)
	f.ServerWithHostName(server, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(server, dist)
	f.ServerWithNotifier(server, notifier)

	// Guards.
	f.ServerWithGuard(server, guards.Session)

	// Pages.
	f.ServerWithPage(server, pages.Welcome)
	f.ServerWithPage(server, pages.Todos)

	// Api.
	f.ServerWithApi(server, api.Events)

	// Start.
	f.ServerStart(server)
}
