package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/controllers/any"
	"main/lib/controllers/events"
	"main/lib/controllers/expired"
	"main/lib/controllers/todos"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	f.NewServer().
		WithAddress("127.0.0.1:8080").
		WithApiController(events.Controller{}).
		WithPageController(any.Controller{}).
		WithPageController(expired.Controller{}).
		WithPageController(todos.Controller{}).
		Start(dist)
}
