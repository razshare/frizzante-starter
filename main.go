package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
	"main/lib/routes"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frizzante.
		NewServer().
		WithNotifier(notifiers.Console).
		WithAddress("127.0.0.1:8080").
		WithRequestHandler("GET /events", routes.GetEvents).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /todos", routes.GetTodos).
		WithRequestHandler("POST /todos", routes.PostTodos).
		WithRequestHandler("GET /welcome", routes.GetWelcome).
		WithRequestHandler("GET /", routes.GetDefault).
		Start(dist)
}
