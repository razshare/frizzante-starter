package main

import (
	"embed"
	"github.com/razshare/frizzante/environments"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/traces"
	"main/lib/routes/handlers"
	"os"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	if err := environments.LoadDotenv(".env"); err != nil {
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
