package main

import (
	"embed"
	"github.com/joho/godotenv"
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/traces"
	"main/lib/handlers"
	"os"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	if err := godotenv.Load(".env"); err != nil {
		traces.Trace(server.ErrorLog, err)
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
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /", Handler: handlers.Default})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /welcome", Handler: handlers.Welcome})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /todos", Handler: handlers.Todos})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /check", Handler: handlers.Check})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /uncheck", Handler: handlers.Uncheck})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /add", Handler: handlers.Add})
	server.Routes = append(server.Routes, servers.Route{Pattern: "GET /remove", Handler: handlers.Remove})

	server.Start()
}
