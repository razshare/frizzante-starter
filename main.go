package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/components/server/events"
	"main/lib/components/server/render"
	"main/lib/components/server/session"
	"main/lib/components/server/todos"
	"main/lib/components/server/welcome"
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
	f.ServerWithGuard(s, render.Guard)
	f.ServerWithGuard(s, session.Guard)

	// Routes.
	f.ServerWithIndex(s, todos.Index)
	f.ServerWithIndex(s, welcome.Index)

	// Api.
	f.ServerWithApi(s, events.Api)

	// Start.
	f.ServerStart(s)
}
