package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/api"
	"main/lib/pages"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	server := f.ServerCreate()
	notifier := f.NotifierCreate()
	archive := f.ArchiveCreateLocal(notifier, ".sessions")
	sessionBuilder := f.SessionBuilderCreate(archive, lib.InitializeState)

	f.ServerWithPort(server, 8080)
	f.ServerWithHostName(server, "127.0.0.1")
	f.ServerWithEmbeddedFileSystem(server, dist)
	f.ServerWithNotifier(server, notifier)
	f.ServerWithApiBuilder(server, api.Events)
	f.ServerWithPageBuilder(server, pages.Todos)
	f.ServerWithPageBuilder(server, pages.Welcome)
	f.ServerWithSessionBuilder[lib.UserSession](server, sessionBuilder)
	f.ServerStart(server)
}
