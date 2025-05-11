package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/api"
	"main/lib/pages"
	"main/lib/sessions"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	// Build notifier.
	notifier := f.NotifierCreate()

	// Build server.
	server := f.ServerCreate()
	f.ServerWithPort(server, 8080)
	f.ServerWithHostName(server, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(server, dist)
	f.ServerWithNotifier(server, notifier)

	// Build sessions.
	f.ServerWithSessionBuilder(server, sessions.Archive)
	//f.ServerWithSessionBuilder(server, sessions.Memory)

	// Build pages.
	f.ServerWithPageBuilder(server, pages.Welcome)
	f.ServerWithPageBuilder(server, pages.Todos)

	// Build api.
	f.ServerWithApiBuilder(server, api.Events)

	// Start.
	f.ServerStart(server)
}
