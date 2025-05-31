package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib/handlers"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	frz.NewServer().
		WithDist(dist).
		WithRequestHandler("GET /", handlers.GetDefault).
		WithRequestHandler("GET /welcome", handlers.GetWelcome).
		WithRequestHandler("GET /todos", handlers.GetTodos).
		WithRequestHandler("POST /todos", handlers.PostTodos).
		Start()
}
