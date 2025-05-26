package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib"
	"main/lib/routes"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frizzante.
		NewServer().
		WithNotifier(lib.Notifier).
		WithAddress("127.0.0.1:8080").
		WithRequestHandler("GET /todos", routes.GetTodos).
		WithRequestHandler("POST /todos", routes.PostTodos).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /events", routes.GetEvents).
		WithRequestHandler("GET /welcome", routes.GetWelcome).
		WithRequestHandler("GET /", routes.GetAny).
		Start(dist)
}
