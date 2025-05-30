package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/handlers"
	"main/lib/notifiers"
)

//go:embed .dist/*/**
var dist embed.FS
var public []frizzante.Guard
var protected = []frizzante.Guard{guards.SessionIsValid}
var server = frizzante.
	NewServer().
	WithNotifier(notifiers.Console).
	WithAddress("127.0.0.1:8080").
	WithRoute("POST /todos", protected, handlers.PostTodos).
	WithRoute("GET /todos", protected, handlers.GetTodos).
	WithRoute("GET /events", protected, handlers.GetEvents).
	WithRoute("GET /expired", public, handlers.GetExpired).
	WithRoute("GET /welcome", public, handlers.GetWelcome).
	WithRoute("GET /", public, handlers.GetDefault)

func main() {
	server.Start(dist)
}
