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
	f.NewServer().
		// Configure.
		WithAddress("127.0.0.1:8080").
		WithEfs(dist).
		// Add page controllers.
		WithPageController(pages.Any).
		WithPageController(pages.Todos).
		WithPageController(pages.Expired).
		// Add api controllers.
		WithApiController(api.Events).
		// Start.
		Start()
}
