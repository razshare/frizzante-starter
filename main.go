package main

import (
	"embed"
	"github.com/razshare/frizzante"
	"main/lib/api"
	"main/lib/notifiers"
)

//go:embed .dist/*/**
var dist embed.FS
var server = frizzante.
	NewServer().
	WithNotifier(notifiers.Console).
	WithAddress("127.0.0.1:8080")

func main() {
	api.Load(server)
	server.Start(dist)
}
