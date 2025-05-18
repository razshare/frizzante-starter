package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/controllers/api"
	"main/lib/controllers/pages"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	// Create.
	server := f.NewServer()
	notifier := f.NewNotifier()

	// Configure.
	server.WithPort(8080)
	server.WithNotifier(notifier)
	server.WithHostName("127.0.0.1")
	server.WithEmbeddedFileSystem(&dist)
	server.WithPageController(pages.WelcomeController{})
	server.WithApiController(api.EventsController{})
	server.WithPageController(pages.TodosController{})

	//Start.
	server.Start()
}
