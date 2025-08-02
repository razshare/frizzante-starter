package main

import (
	"embed"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"main/lib/routes/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs
	server.Addr = "127.0.0.1:7777"

	server.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handlers.Default},
		{Pattern: "GET /welcome", Handler: handlers.Welcome},
		{Pattern: "GET /todos", Handler: handlers.Todos},
		{Pattern: "GET /check", Handler: handlers.Check},
		{Pattern: "GET /uncheck", Handler: handlers.Uncheck},
		{Pattern: "GET /add", Handler: handlers.Add},
		{Pattern: "GET /remove", Handler: handlers.Remove},
	}

	server.Start()
}
