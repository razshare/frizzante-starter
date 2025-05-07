package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/api"
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

	// Build pages.
	f.ServerWithPageBuilder(server, pages.Welcome)
	f.ServerWithPageBuilder(server, pages.Todos)

	// Build api.
	f.ServerWithApiBuilder(server, api.Events)

	// Start.
	f.ServerStart(server)
}
