package main

import (
	"embed"
	"github.com/joho/godotenv"
	"main/lib/handlers"
	"os"

	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	if err := godotenv.Load(".env"); err != nil {
		server.Notifier.SendError(err)
	} else {
		server.Address = os.Getenv("server.address")
		server.SecureAddress = os.Getenv("server.secure_address")
		server.Key = os.Getenv("server.key")
		server.Certificate = os.Getenv("server.certificate")
		server.PublicRoot = os.Getenv("server.public_root")
		server.AppRoot = os.Getenv("server.app_root")
		server.ServerJs = os.Getenv("server.server_js")
		server.IndexHtml = os.Getenv("server.index_html")
	}

	server.Efs = efs
	server.AddRoute(routes.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(routes.Route{Pattern: "GET /welcome", Handler: handlers.Welcome})
	server.AddRoute(routes.Route{Pattern: "GET /todos", Handler: handlers.Todos})
	server.AddRoute(routes.Route{Pattern: "GET /check", Handler: handlers.Check})
	server.AddRoute(routes.Route{Pattern: "GET /uncheck", Handler: handlers.Uncheck})
	server.AddRoute(routes.Route{Pattern: "GET /add", Handler: handlers.Add})
	server.AddRoute(routes.Route{Pattern: "GET /remove", Handler: handlers.Remove})
	server.Start()
}
