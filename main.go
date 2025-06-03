package main

import (
	"embed"
	"github.com/razshare/frizzante/frz"
	"main/lib/handlers"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frz.NewServer().
		WithDist(dist).
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.GetDefault}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.GetWelcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.GetTodos}).
		AddRoute(frz.Route{Pattern: "POST /todos", Handler: handlers.PostTodos}).
		Start()
}
