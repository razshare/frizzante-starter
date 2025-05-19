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

	// Pages.
	server.WithPageController(pages.WelcomeController{})
	server.WithPageController(pages.TodosController{})
	server.WithPageController(pages.ExpiredController{})

	// Api.
	server.WithApiController(api.EventsController{})

	//Start.
	server.Start()
}
