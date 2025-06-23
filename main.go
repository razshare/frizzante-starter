package main

import (
	"embed"
	"github.com/razshare/frizzante/libsrv"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = libsrv.NewServer()

func main() {
	server.Efs = efs
	server.AddRoute(libsrv.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(libsrv.Route{Pattern: "GET /welcome", Handler: handlers.Welcome})
	server.AddRoute(libsrv.Route{Pattern: "GET /todos", Handler: handlers.Todos})
	server.AddRoute(libsrv.Route{Pattern: "GET /check", Handler: handlers.Check})
	server.AddRoute(libsrv.Route{Pattern: "GET /uncheck", Handler: handlers.Uncheck})
	server.AddRoute(libsrv.Route{Pattern: "GET /add", Handler: handlers.Add})
	server.AddRoute(libsrv.Route{Pattern: "GET /remove", Handler: handlers.Remove})
	server.Start()
}
