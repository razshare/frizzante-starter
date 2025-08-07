package main

import (
	"embed"
	"github.com/razshare/frizzante/server"
	"main/lib/routes/handlers"
)

//go:embed app/dist
var efs embed.FS
var srv = server.Default()

func main() {
	srv.Efs = efs

	srv.Routes = []server.Route{
		{Pattern: "GET /", Handler: handlers.Default},
		{Pattern: "GET /welcome", Handler: handlers.Welcome},
		{Pattern: "GET /todos", Handler: handlers.Todos},
		{Pattern: "GET /check", Handler: handlers.Check},
		{Pattern: "GET /uncheck", Handler: handlers.Uncheck},
		{Pattern: "GET /add", Handler: handlers.Add},
		{Pattern: "GET /remove", Handler: handlers.Remove},
	}

	server.Start(srv)
}
