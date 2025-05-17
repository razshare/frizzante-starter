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
	f.ServerWithNotifier(server, notifier)
	f.ServerWithHostName(server, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(server, &dist)

	//Pages.
	f.ServerWithPageBuilder(server, pages.Todos)
	f.ServerWithPageBuilder(server, pages.Welcome)
	f.ServerWithPageBuilder(server, pages.Expired)

	// Api.
	f.ServerWithApiBuilder(server, api.Events)

	//Start.
	f.ServerStart(server)
}
