package main

import (
	"embed"
	frz "github.com/razshare/frizzante"
	"main/lib/handlers"
)

//go:embed .dist/*/**
var dist embed.FS
var public []frz.Guard

func main() {
	frz.NewServer().
		WithDist(dist).
		Map(public, "GET /", handlers.GetDefault).
		Map(public, "GET /welcome", handlers.GetWelcome).
		Map(public, "GET /todos", handlers.GetTodos).
		Map(public, "POST /todos", handlers.PostTodos).
		Start()
}
